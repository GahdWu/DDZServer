package rpcServer

// 响应对象
type ResponseObject struct {
	// 响应结果的状态值
	Code *ResultStatus

	// 响应结果的状态值所对应的描述信息
	Message string

	// 响应结果的数据
	Data interface{}
}

// 设置响应结果的状态值
// rs：响应结果的状态值
// 返回值：
// 响应结果对象
func (responseObj ResponseObject) SetResultStatus(rs *ResultStatus) ResponseObject {
	responseObj.Code = rs
	responseObj.Message = rs.Message

	return responseObj
}

// 设置响应结果的数据
// data：响应结果的数据
// 返回值：无
func (responseObj *ResponseObject) SetData(data interface{}) {
	responseObj.Data = data
}

// 获取初始的响应对象
// 返回值：
// 响应对象
func NewResponseObj() *ResponseObject {
	return &ResponseObject{
		Code:    Success,
		Message: "",
		Data:    nil,
	}
}
