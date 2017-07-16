package rpcServer

import (
	"encoding/json"
	"fmt"
)

// 发送响应结果
// clientObj：客户端对象
// requestObj：请求对象（如果为nil则代表是服务端主动推送信息，否则为客户端请求信息）
// responseObject：响应对象（不能为指针类型，否则在registerFunction时判断类型会出错）
// priority:优先级
func ResponseResult(clientObj *Client, requestObj *requestObject, responseObj interface{}, priority Priority) {
	message, err := json.Marshal(responseObj)
	if err != nil {
		clientObj.WriteLog(fmt.Sprintf("序列化输出结果%v出错", responseObj))
		return
	}

	var sendDataItemObj *sendDataItem

	if requestObj == nil {
		sendDataItemObj = newSendDataItem(0, message)
	} else {
		sendDataItemObj = newSendDataItem(requestObj.Id, message)
	}

	clientObj.appendSendData(sendDataItemObj, priority)
}
