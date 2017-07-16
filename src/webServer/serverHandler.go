package webServer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Gahd/DDZServer/src/config"
	. "github.com/Gahd/DDZServer/src/model/responseObject"
	"moqikaka.com/goutil/debugUtil"
	"moqikaka.com/goutil/logUtil"
	"moqikaka.com/goutil/webUtil"
)

// 定义服务处理对象
type ServerHandler struct {
}

//服务处理逻辑
func (this *ServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//过滤请求
	if isNeedFilterRequest(w, r) {
		return
	}

	//创建返回对象
	result := NewResponseObject()

	startTime := time.Now().Unix()
	endTime := time.Now().Unix()

	if debugUtil.IsDebug() {
		//允许跨域
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}

	// 在输出结果给客户端之后再来处理日志的记录，以便于可以尽快地返回给客户端
	defer func() {
		// 为了避免在调用ParseForm之前就提前返回，所以调用ParseForm方法
		r.ParseForm()

		// 获取输入参数的字符串形式
		parameter := ""
		if len(r.Form) > 0 {
			parameter_byte, _ := json.Marshal(r.Form)
			parameter = string(parameter_byte)
		}

		// 记录DEBUG日志
		useSeconds := endTime - startTime
		if debugUtil.IsDebug() || result.ResultStatus != Success || useSeconds > 3 {
			resultData, _ := json.Marshal(result)

			msg := fmt.Sprintf("IP:%s->%s 请求数据:%v 返回数据:%s 用时:%vs", webUtil.GetRequestIP(r), r.RequestURI, parameter, string(resultData), useSeconds)
			logUtil.NormalLog(msg, logUtil.Debug)
			debugUtil.Println(msg)
		}
	}()

	//获取业务处理器
	handler := GetHandler(r.RequestURI)
	if handler == nil {
		//未找到接口
		result.SetResultStatus(APINotDefined)
		result.SetData(config.GetMonitorConfig().MonitorServerName)

		//输出结果
		writeResult(w, result)
		return
	}

	// 验证IP
	if rs := handler.CheckIP(r); rs != Success {
		writeResult(w, result.SetResultStatus(rs))
		return
	}

	// 检查参数
	if rs := handler.CheckParam(r); rs != Success {
		writeResult(w, result.SetResultStatus(rs))
		return
	}

	//处理数据
	result = handler.HandlerFunc(w, r)
	endTime = time.Now().Unix()

	//输出结果
	writeResult(w, result)
}

//过滤请求
func isNeedFilterRequest(w http.ResponseWriter, r *http.Request) bool {
	switch r.RequestURI {
	case "/": //首页，测试
		w.Write([]byte("OK"))
		return true
	case "/favicon.ico": //图标请求，不返回任何数据
		return true
	}

	return false
}

//输出结果
func writeResult(w http.ResponseWriter, result *ResponseObject) {
	//转换为json输出
	if dataBytes, err := json.Marshal(result); err == nil {
		//发送数据
		w.Write(dataBytes)
	}
}
