package rpcServer

import (
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"github.com/Gahd/DDZServer/src/redisMgr"
	"moqikaka.com/goutil/fileUtil"
	"moqikaka.com/goutil/intAndBytesUtil"
	"moqikaka.com/goutil/logUtil"
	"moqikaka.com/goutil/timeUtil"
	"moqikaka.com/goutil/zlibUtil"
	"net"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const (
	// 包头的长度
	con_HEADER_LENGTH = 4

	// 定义请求、响应数据的前缀的长度
	con_ID_LENGTH = 4

	// 客户端失效的秒数
	con_CLIENT_EXPIRE_SECONDS = 300 * time.Second
)

var (
	// 字节的大小端顺序
	byterOrder = binary.LittleEndian

	// 全局客户端的id，从1开始进行自增
	globalClientId int32 = 0
)

// 定义客户端对象，以实现对客户端连接的封装
type Client struct {
	// 唯一标识
	id int32

	// 客户端连接对象
	conn net.Conn

	//连接状态 (1:连接中, 2:最后一条消息，3,断开)
	connStatus ConnStatus

	// 接收到的消息内容
	receiveData []byte

	// 待发送的数据
	sendData []*sendDataItem

	// 低优先级的待发送的数据
	sendData_LowPriority []*sendDataItem

	// 锁对象（用于控制对sendDatap、sendData_LowPriority的并发访问；receiveData不需要，因为是同步访问）
	mutex sync.Mutex

	// 上次活跃时间
	activeTime time.Time

	// 玩家id
	playerId string

	// 玩家当前正在处理的方法
	curHandleFunc string
}

// 设置玩家正在处理的方法
func (clientObj *Client) setCurFuncName(value string) {
	clientObj.curHandleFunc = value
}

// 获取玩家正在处理的方法
func (clientObj *Client) GetCurFuncName() string {
	return clientObj.curHandleFunc
}

// 获取连接状态
func (clientObj *Client) GetConnStatus() ConnStatus {
	return clientObj.connStatus
}

// 设置连接状态
func (clientObj *Client) setConnStatus(status ConnStatus) {
	clientObj.connStatus = status
}

// 获取远程地址（IP_Port）
func (clientObj *Client) GetRemoteAddr() string {
	items := strings.Split(clientObj.conn.RemoteAddr().String(), ":")

	return fmt.Sprintf("%s_%s", items[0], items[1])
}

// 获取远程地址（IP）
func (clientObj *Client) GetRemoteShortAddr() string {
	items := strings.Split(clientObj.conn.RemoteAddr().String(), ":")

	return items[0]
}

// 获取client的Id
// 返回值：
// Id
func (clientObj *Client) GetId() int32 {
	return clientObj.id
}

// 获取playerId
// 返回值：
// playerId
func (clientObj *Client) GetPlayerId() string {
	return clientObj.playerId
}

// 获取接收到的数据
// 返回值：
// 消息对应客户端的唯一标识
// 消息内容
// 是否含有有效数据
func (clientObj *Client) getReceiveData() (id int32, message []byte, exists bool) {

	// 判断是否包含头部信息
	if len(clientObj.receiveData) < con_HEADER_LENGTH {
		return
	}

	// 获取头部信息
	header := clientObj.receiveData[:con_HEADER_LENGTH]

	// 将头部数据转换为内容的长度
	contentLength := intAndBytesUtil.BytesToInt32(header, byterOrder)

	// 处理内部未处理的异常，以免导致玩家协程崩溃
	defer func() {
		if r := recover(); r != nil {
			logUtil.LogUnknownError(r)
			logUtil.Log(fmt.Sprintf("函数getReceiveData中发生崩溃， 玩家：%v, len(clientObj.receiveData):%d, header:%d, con_HEADER_LENGTH:%d,contentLength:%d",
				clientObj.GetPlayerId(), len(clientObj.receiveData), header, con_HEADER_LENGTH, contentLength), logUtil.Error, true)
			panic("函数getReceiveData中接收数据发生崩溃")
		}
	}()

	// 判断长度是否满足
	if len(clientObj.receiveData) < con_HEADER_LENGTH+int(contentLength) {
		return
	}

	// 存在有效的消息
	exists = true

	// 提取消息内容
	content := clientObj.receiveData[con_HEADER_LENGTH : con_HEADER_LENGTH+contentLength]

	// 将对应的数据截断，以得到新的内容
	clientObj.receiveData = clientObj.receiveData[con_HEADER_LENGTH+contentLength:]

	// 判断是否为心跳包，如果是心跳包，则不解析，直接返回
	if contentLength == 0 || len(content) == 0 {
		return
	}

	// 判断内容的长度是否足够
	if len(content) < con_ID_LENGTH {
		logUtil.Log(fmt.Sprintf("内容数据不正确；con_ID_LENGTH=%d,len(content)=%d", con_ID_LENGTH, len(content)), logUtil.Warn, true)
	}

	// 将内容分隔为2部分
	idBytes, content := content[:con_ID_LENGTH], content[con_ID_LENGTH:]

	// 提取id、message
	id = intAndBytesUtil.BytesToInt32(idBytes, byterOrder)
	message = content

	clientObj.WriteLog(fmt.Sprintf("收到消息:%d:%s", id, string(content)))

	return
}

// 获取待发送的数据
// 返回值：
// 待发送数据项
// 是否含有有效数据
func (clientObj *Client) getSendData() (sendDataItemObj *sendDataItem, exists bool) {
	clientObj.mutex.Lock()
	defer clientObj.mutex.Unlock()

	// 如果没有数据则直接返回
	if len(clientObj.sendData) == 0 {
		return
	}

	// 取出第一条数据,并为返回值赋值
	sendDataItemObj = clientObj.sendData[0]
	exists = true

	// 删除已经取出的数据
	clientObj.sendData = clientObj.sendData[1:]

	return
}

// 获取低优先级待发送的数据
// 返回值：
// 待发送数据项
// 是否含有有效数据
func (clientObj *Client) getSendData_LowPriority() (sendDataItemObj *sendDataItem, exists bool) {
	clientObj.mutex.Lock()
	defer clientObj.mutex.Unlock()

	// 如果没有数据则直接返回
	if len(clientObj.sendData_LowPriority) == 0 {
		return
	}

	// 取出第一条数据,并为返回值赋值
	sendDataItemObj = clientObj.sendData_LowPriority[0]
	exists = true

	// 删除已经取出的数据
	clientObj.sendData_LowPriority = clientObj.sendData_LowPriority[1:]

	return
}

// 追加接收到的数据
// receiveData：接收到的数据
// 返回值：无
func (clientObj *Client) appendReceiveData(receiveData []byte) {
	clientObj.receiveData = append(clientObj.receiveData, receiveData...)
	clientObj.activeTime = time.Now()
}

// 追加发送的数据
// sendDataItemObj:待发送数据项
// priority:优先级
// 返回值：无
func (clientObj *Client) appendSendData(sendDataItemObj *sendDataItem, priority Priority) {
	clientObj.mutex.Lock()
	defer clientObj.mutex.Unlock()

	if priority == Con_LowPriority {
		clientObj.sendData_LowPriority = append(clientObj.sendData_LowPriority, sendDataItemObj)
	} else {
		clientObj.sendData = append(clientObj.sendData, sendDataItemObj)
	}
}

// 判断客户端是否超时
// 返回值：是否超时
func (clientObj *Client) expired() bool {
	return time.Now().Unix() > clientObj.activeTime.Add(con_CLIENT_EXPIRE_SECONDS).Unix()
}

// 发送消息
// id：需要添加到内容前的数据
// sendDataItemObj：待发送数据项
// 返回值：
// 错误对象
func (clientObj *Client) sendMessage(sendDataItemObj *sendDataItem) error {
	clientObj.WriteLog(fmt.Sprintf("发送消息,%d:%s", sendDataItemObj.id, string(sendDataItemObj.data)))

	beforeTime := time.Now().Unix()
	afterTime := time.Now().Unix()
	defer func() {
		// 如果发送的时间超过1秒，则记录下来
		if afterTime-beforeTime > 1 {
			redisMgr.RecordExpireLog(fmt.Sprintf("id:%d, size:%d, UseTime:%d", sendDataItemObj.id, len(sendDataItemObj.data), afterTime-beforeTime))
		}
	}()

	idBytes := intAndBytesUtil.Int32ToBytes(sendDataItemObj.id, byterOrder)

	// 将idByte和内容合并
	content := append(idBytes, sendDataItemObj.data...)

	// 获得数组的长度
	contentLength := len(content)

	// 将长度转化为字节数组
	header := intAndBytesUtil.Int32ToBytes(int32(contentLength), byterOrder)

	// 将头部与内容组合在一起
	message := append(header, content...)

	// if config.IfCompress {
	if false {
		//将数据进行zlib压缩
		message, err := zlibUtil.Compress(message, zlib.DefaultCompression)
		if err != nil {
			logUtil.Log(fmt.Sprintf("函数sendMessage,压缩message出错,message:%v，错误信息为：%s", message, err), logUtil.Error, true)
			return err
		}
	}

	// 发送消息
	_, err := clientObj.conn.Write(message)
	if err != nil {
		clientObj.WriteLog(fmt.Sprintf("发送消息,%d:%s,出现错误：%s", sendDataItemObj.id, string(sendDataItemObj.data), err))
	}

	afterTime = time.Now().Unix()

	return err
}

// 玩家登陆
// playerId：玩家id
// 返回值：无
func (clientObj *Client) PlayerLogin(playerId string) {
	clientObj.playerId = playerId

	playerLogin(clientObj)
}

// 玩家登出，客户端退出
// msg：退出时受到的消息
// 返回值：无
func (clientObj *Client) LogoutAndQuit(msg string) {
	clientObj.WriteLog(fmt.Sprintf("收到退出消息:%s", msg))
	clientObj.playerLogout()
	clientObj.quit()
}

// 保护性退出,保证所有数据推送成功
// msg：退出时受到的消息
// 返回值：
// 是否退出成功
func (clientObj *Client) ProtectLogoutAndQuit(msg string) {
	clientObj.mutex.Lock()
	defer clientObj.mutex.Unlock()

	//设置conn状态为2
	if clientObj.connStatus == con_Open {
		clientObj.setConnStatus(con_WaitForClose)
	}
}

// 返回值：无
func (clientObj *Client) playerLogout() {
	deletePlayer(clientObj.playerId)

	clientObj.playerId = ""
}

// 退出
// 返回值：无
func (clientObj *Client) quit() {
	clientObj.conn.Close()
	clientObj.setConnStatus(con_Close)
}

// 格式化
func (clientObj *Client) String() string {
	return fmt.Sprintf("{Id:%d, RemoteAddr:%d, activeTime:%s, playerId:%s}", clientObj.id, clientObj.GetRemoteAddr(), timeUtil.Format(clientObj.activeTime, "yyyy-MM-dd HH:mm:ss"), clientObj.playerId)
}

// 记录日志
// log：日志内容
func (clientObj *Client) WriteLog(log string) {
	if debug {
		fileUtil.WriteFile("Log", clientObj.GetRemoteAddr(), true,
			timeUtil.Format(time.Now(), "yyyy-MM-dd HH:mm:ss"),
			" ",
			fmt.Sprintf("client:%s", clientObj.String()),
			" ",
			log,
			"\r\n",
			"\r\n",
		)
		if clientObj.playerId != "" {
			fileUtil.WriteFile("Log", clientObj.playerId, true,
				timeUtil.Format(time.Now(), "yyyy-MM-dd HH:mm:ss"),
				" ",
				fmt.Sprintf("client:%s", clientObj.String()),
				" ",
				log,
				"\r\n",
				"\r\n",
			)
		}
	}
}

// 新建客户端对象
// conn：连接对象
// 返回值：客户端对象的指针
func newClient(_conn net.Conn) *Client {
	// 生成Id的方法
	generateId := func() int32 {
		atomic.AddInt32(&globalClientId, 1)

		return globalClientId
	}

	// 创建新的客户端对象
	clientObj := &Client{
		id:                   generateId(),
		conn:                 _conn,
		connStatus:           con_Open,
		receiveData:          make([]byte, 0, 1024),
		sendData:             make([]*sendDataItem, 0, 16),
		sendData_LowPriority: make([]*sendDataItem, 0, 16),
		activeTime:           time.Now(),
		playerId:             "",
		// 与玩家相关的属性使用默认值
	}

	return clientObj
}
