package redisMgr

import (
	"fmt"
	"moqikaka.com/goutil/logUtil"
)

var (
	//是否关闭超时日志
	IsCloseExpireLog = false
)

// 记录超时日志
// log：日志信息
// 返回值：无
func RecordExpireLog(log string) {
	if IsCloseExpireLog {
		return
	}

	conn := getRedisConn()
	defer conn.Close()

	// 再将排行数据添加到列表中
	valueList := make([]interface{}, 0, 2)
	valueList = append(valueList, con_ExpireLog_Key)
	valueList = append(valueList, log)

	if _, err := conn.Do("RPUSH", valueList...); err != nil {
		logUtil.Log(fmt.Sprintf("将%s超时数据保存到Redis中出错，错误信息为：%s", log, err), logUtil.Error, true)
	}
}
