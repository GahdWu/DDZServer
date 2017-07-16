package rpcServer

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"moqikaka.com/goutil/logUtil"
)

var (
	// 客户端、玩家集合
	clientMap = make(map[int32]*Client)
	playerMap = make(map[string]*Client)

	// 读写锁
	mutexForClient sync.RWMutex
	mutexForPlayer sync.RWMutex
)

const (
	// 玩家登陆、重登录的模块名称
	con_PlayerLoginModuleName = "Player"

	// 玩家登陆的方法名称
	con_PlayerLoginMethodName = "Login"

	// 玩家重新登录的方法名称
	con_PlayerReloginMethodName = "Relogin"
)

func init() {
	// 处理client过期
	go func() {
		// 处理内部未处理的异常，以免导致主线程退出，从而导致系统崩溃
		defer func() {
			if r := recover(); r != nil {
				logUtil.LogUnknownError(r)
			}
		}()

		for {
			// 因为刚开始时不存在过期，所以先暂停5分钟
			time.Sleep(5 * time.Minute)

			// 获取客户端连接列表
			clientList := getClientList()

			// 记录日志
			logUtil.Log(fmt.Sprintf("当前客户端数量为：%d", len(clientList)), logUtil.Debug, true)

			for _, clientObj := range clientList {
				if clientObj.expired() {
					clientObj.LogoutAndQuit("因为不活跃被清理掉")
				}
			}
		}
	}()
}

// 获取客户端列表
// 返回值：
// 客户端列表
func getClientList() (clientList []*Client) {
	mutexForClient.RLock()
	defer mutexForClient.RUnlock()

	for _, value := range clientMap {
		clientList = append(clientList, value)
	}

	return
}

//获取客户端和player数量
//返回值:
//客户端数量, 玩家数量
func GetClientAndPlayerNum() (int, int) {
	// 获取客户端数量
	clientNum := func() int {
		mutexForClient.RLock()
		defer mutexForClient.RUnlock()

		return len(clientMap)
	}

	// 获取玩家数量
	playerNum := func() int {
		mutexForPlayer.RLock()
		defer mutexForPlayer.RUnlock()

		return len(playerMap)
	}

	return clientNum(), playerNum()
}

// 根据玩家对象获取对应的客户端对象
// id：客户端Id
// 返回值：
// 客户端对象
// 是否存在客户端对象
func GetClientById(id int32) (*Client, bool) {
	mutexForClient.RLock()
	defer mutexForClient.RUnlock()

	if clientObj, exists := clientMap[id]; exists {
		return clientObj, true
	}

	return nil, false
}

// 根据玩家Id获得客户端对象
// playerId：玩家Id
// 返回值：
// 客户端对象
// 是否存在客户端对象
func GetClientByPlayerId(playerId string) (*Client, bool) {
	mutexForPlayer.RLock()
	defer mutexForPlayer.RUnlock()

	if clientObj, exists := playerMap[playerId]; exists {
		return clientObj, true
	}

	return nil, false
}

// 添加新的客户端
// clientObj：客户端对象
func registerClient(clientObj *Client) {
	mutexForClient.Lock()
	defer mutexForClient.Unlock()

	clientMap[clientObj.id] = clientObj
}

// 移除客户端
// clientObj：客户端对象
func unRegisterClient(clientObj *Client) {
	// 先删除玩家
	deletePlayer(clientObj.playerId)

	// 再删除客户端对象
	mutexForClient.Lock()
	defer mutexForClient.Unlock()

	delete(clientMap, clientObj.id)
}

// 玩家登陆
// clientObj：客户端对象
func playerLogin(clientObj *Client) {
	mutexForPlayer.Lock()
	defer mutexForPlayer.Unlock()

	playerMap[clientObj.playerId] = clientObj
}

// 删除玩家
// playerId：玩家Id
func deletePlayer(playerId string) {
	if playerId == "" {
		return
	}

	mutexForPlayer.Lock()
	defer mutexForPlayer.Unlock()

	delete(playerMap, playerId)
}

// 处理请求
// clientObj：对应的客户端对象
// id：客户端请求唯一标识
// request：请求内容字节数组(json格式)
// 返回值：无
func handleRequest(clientObj *Client, id int32, request []byte) {
	responseObj := NewResponseObj()

	// 解析请求字符串
	var requestObj requestObject

	// 提取请求内容
	if err := json.Unmarshal(request, &requestObj); err != nil {
		clientObj.WriteLog(fmt.Sprintf("反序列化%s出错，错误信息为：%s", string(request), err))
		ResponseResult(clientObj, &requestObj, responseObj.SetResultStatus(DataFormatError), Con_HighPriority)
		return
	}

	// 对requestObj的属性Id赋值
	requestObj.Id = id

	// 对参数要特殊处理：将Player或Client对象作为第一个参数(如果是登录或重新登录，是Client对象，否则是Player对象)
	parameters := make([]interface{}, 0)
	if requestObj.ModuleName == con_PlayerLoginModuleName &&
		(requestObj.MethodName == con_PlayerLoginMethodName || requestObj.MethodName == con_PlayerReloginMethodName) {
		parameters = append(parameters, interface{}(clientObj))
		parameters = append(parameters, requestObj.Parameters...)
	} else {
		// 判断玩家是否已经登陆
		if clientObj.playerId == "" {
			ResponseResult(clientObj, &requestObj, responseObj.SetResultStatus(PlayerNotLogin), Con_HighPriority)
			return
		}

		// 判断是否能找到玩家
		playerObj, rs := getPlayerFunc(clientObj, clientObj.playerId)
		if rs != Success {
			ResponseResult(clientObj, &requestObj, responseObj.SetResultStatus(rs), Con_HighPriority)
			return
		}

		parameters = append(parameters, interface{}(playerObj))
		parameters = append(parameters, requestObj.Parameters...)
	}

	// 为参数赋值
	requestObj.Parameters = parameters

	// 调用方法
	callFunction(clientObj, &requestObj)
}
