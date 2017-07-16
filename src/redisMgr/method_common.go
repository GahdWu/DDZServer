package redisMgr

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
	"moqikaka.com/goutil/logUtil"
)

// 设置玩家当前所在模块
// playerId:玩家id
// 返回值
// int:模块id
// error:错误
func GetPlayerCurModuleType(playerId string) (int, error) {
	conn := getRedisConn()
	defer conn.Close()

	reply, err := conn.Do("HGET", getPlayerKey(playerId), con_CurModuleType_field)
	if err != nil {
		if err == redis.ErrNil {
			return 0, nil
		} else {
			logUtil.Log(fmt.Sprintf("从Redis中获取玩家复活结束时间出错1，错误信息为：%s", err), logUtil.Error, true)
			return 0, err
		}
	}

	curModule, err := redis.Int(reply, err)
	if err != nil {
		if err == redis.ErrNil {
			return 0, nil
		} else {
			logUtil.Log(fmt.Sprintf("从Redis中获取玩家复活结束时间出错2，错误信息为：%s", err), logUtil.Error, true)
			return 0, err
		}
	}

	return curModule, nil
}

//获取服务器ip
// conn:Redis连接
// playerKey：玩家Key
func getSocketServerAddressByKey(conn redis.Conn, playerKey string) (string, error) {
	reply, err := conn.Do("HGET", playerKey, con_SocketServerAddress_field)
	if err != nil {
		logUtil.Log(fmt.Sprintf("获取socketServerAddress出错1，错误信息为：%s", err), logUtil.Error, true)
		return "", err
	}

	address, err := redis.String(reply, err)
	if err != nil {
		if err == redis.ErrNil {
			return "", err
		} else {
			logUtil.Log(fmt.Sprintf("获取socketServerAddress出错2，错误信息为：%s", err), logUtil.Error, true)
			return "", err
		}
	}

	return address, nil
}

//// 设置玩家状态信息
//// conn：redis连接
//// playerKey：玩家Id
//// status：玩家状态
//func setPlayerStatusByKey(conn redis.Conn, playerKey string, status playerStatus.PlayerStatus) error {
//	_, err := conn.Do("HSET", playerKey, con_Status_field, status)
//	if err != nil {
//		logUtil.Log(fmt.Sprintf("设置玩家状态出错，错误信息为：%s", err), logUtil.Error, true)
//		return err
//	}

//	return nil
//}

//遍历redis中的所有玩家,检测状态是否正确,不正确就重置
//func CheckRedisPlayer() error {
//	conn := getRedisConn()
//	defer conn.Close()

//	reply, err := conn.Do("KEYS", getAllPlayerKeysPattern())
//	if err != nil {
//		logUtil.Log(fmt.Sprintf("从Redis中判断玩家Id出错1，错误信息为：%s", err), logUtil.Error, true)
//		return err
//	}

//	result, err := redis.Strings(reply, err)
//	if err != nil {
//		logUtil.Log(fmt.Sprintf("从Redis中判断玩家Id出错2，错误信息为：%s", err), logUtil.Error, true)
//		return err
//	}

//	//遍历
//	for _, playerInfo := range result {
//		playerKey := string(playerInfo)

//		//获取玩家状态
//		status, err := getPlayerStatusByKey(conn, playerKey)
//		if err != nil {
//			if err == redis.ErrNil {
//				continue
//			}
//			return err
//		}

//		//获取玩家所在服务器address
//		address, err := getSocketServerAddressByKey(conn, playerKey)
//		if err != nil {
//			if err == redis.ErrNil {
//				continue
//			}
//			return err
//		}

//		//判断是否修改状态
//		if status != playerStatus.Con_End && address == config.GetRpcConfig().GetSocketServerPublishAddress() {
//			err := setPlayerStatusByKey(conn, playerKey, playerStatus.Con_End)
//			if err != nil {
//				return err
//			}
//		}
//	}

//	return nil
//}

// 根据玩家id获取玩家状态
// playerId: 玩家id
// 返回值:
//
//func GetPlayerStatusByPlayerId(playerId string) (playerStatus.PlayerStatus, error) {
//	conn := getRedisConn()
//	defer conn.Close()

//	reply, err := conn.Do("HGET", getPlayerKey(playerId), con_Status_field)
//	if err != nil {
//		logUtil.Log(fmt.Sprintf("获取玩家状态出错1，错误信息为：%s", err), logUtil.Error, true)
//		return -1, err
//	}

//	status, err := redis.Int(reply, err)
//	if err != nil {
//		if err == redis.ErrNil {
//			return -1, err
//		} else {
//			logUtil.Log(fmt.Sprintf("获取玩家状态出错2，错误信息为：%s", err), logUtil.Error, true)
//			return -1, err
//		}
//	}

//	if config.DEBUG {
//		logUtil.Log(fmt.Sprintf("获取玩家状态成功,状态为：%s", status), logUtil.Info, true)
//	}

//	return playerStatus.PlayerStatus(status), nil
//}

// 传入playerId,判断redis中是否存在
// playerId：玩家Id
// 返回值：
// 是否存在
// 错误对象
func CheckPlayerIdIsExist(playerId string) (bool, error) {
	conn := getRedisConn()
	defer conn.Close()

	reply, err := conn.Do("EXISTS", getPlayerKey(playerId))
	if err != nil {
		logUtil.Log(fmt.Sprintf("从Redis中判断玩家Id出错1，错误信息为：%s", err), logUtil.Error, true)
		return false, err
	}

	result, err := redis.Bool(reply, err)
	if err != nil {
		logUtil.Log(fmt.Sprintf("从Redis中判断玩家Id出错2，错误信息为：%s", err), logUtil.Error, true)
		return false, err
	}

	return result, nil
}

// 根据玩家完整Key获取玩家状态
// conn:Redis连接
// playerKey：玩家Key
//func getPlayerStatusByKey(conn redis.Conn, playerKey string) (playerStatus.PlayerStatus, error) {
//	reply, err := conn.Do("HGET", playerKey, con_Status_field)
//	if err != nil {
//		logUtil.Log(fmt.Sprintf("在函数getPlayerStatusByKey中,获取玩家状态出错1，错误信息为：%s", err), logUtil.Error, true)
//		return -1, err
//	}

//	status, err := redis.Int(reply, err)
//	if err != nil {
//		if err == redis.ErrNil {
//			return -1, err
//		} else {
//			logUtil.Log(fmt.Sprintf("在函数getPlayerStatusByKey中,获取玩家状态出错2，错误信息为：%s", err), logUtil.Error, true)
//			return -1, err
//		}
//	}

//	if config.DEBUG {
//		logUtil.Log(fmt.Sprintf("在函数getPlayerStatusByKey中,获取玩家状态成功,状态为：%s", status), logUtil.Info, true)
//	}

//	return playerStatus.PlayerStatus(status), nil
//}
