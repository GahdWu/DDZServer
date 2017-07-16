package redisMgr

import (
	"errors"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
	"moqikaka.com/goutil/debugUtil"
	"moqikaka.com/goutil/logUtil"
)

// 启动Redis连接池
// redisConnection：redis连接字符串
// password：redis连接密码
// database：Redis服务器选择的数据库
// maxActive：redis连接池最大活跃数量
// maxIdle：redis连接池最大空闲数量
// idleTimeout：连接被回收前的空闲时间
// dialConnectTimeout：连接Redis服务器超时时间
// centerAddress：SocketServerCenter地址
// centerAPIAddress : 提供Web API的地址
// 返回值：无
func Start(redisConn, passWord string, database, maxActive, maxIdle int, idleTimeout, dialConnectTimeout time.Duration) {
	debugUtil.Println("redisConfig parameters:", redisConn, passWord, database, maxActive, maxIdle, idleTimeout, dialConnectTimeout)

	// 创建redis连接池
	createRedisPool(redisConn, passWord, database, maxActive, maxIdle, idleTimeout, dialConnectTimeout)

	// 从redis中获取socketServerCenter地址并保存
	//	initSocketServerCenterAddr()

	// 通知初始化成功
	notifyInitSuccess()

	//检测redis
	//	err := CheckRedisPlayer()
	//	if err != nil {
	//		panic("检测redis中的玩家状态出错")
	//	}
}

// 从redis中获取socketServerCenter地址并保存
// 返回值：无
func initSocketServerCenterAddr() {
	conn := getRedisConn()
	defer conn.Close()

	reply, err := conn.Do("GET", con_SocketServerCenterAddress_Key)
	if err != nil {
		errInfo := fmt.Sprintf("从redis获取socketServerCenterAddress地址失败，错误信息为：%s", err)
		logUtil.Log(errInfo, logUtil.Error, true)
		panic(errors.New(errInfo))
	}

	//保存socketServerCenter地址
	socketServerCenterAddress, err = redis.String(reply, err)
	if err != nil {
		errInfo := fmt.Sprintf("字符串化address失败，错误信息为：%s", err)
		logUtil.Log(errInfo, logUtil.Error, true)
		panic(errors.New(errInfo))
	}

	debugUtil.Println("socketServerCenterAddress:", socketServerCenterAddress)
}
