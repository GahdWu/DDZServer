package webServer

import (
	"net/http"

	. "github.com/Gahd/DDZServer/src/model/responseObject"
	"moqikaka.com/Framework/ipMgr"
	"moqikaka.com/goutil/debugUtil"
	"moqikaka.com/goutil/webUtil"
)

// 请求方法对象
type Handler struct {
	// 注册的访问路径
	Path string

	// 方法定义
	HandlerFunc func(http.ResponseWriter, *http.Request) *ResponseObject

	// 是否需要验证IP
	IsCheckIP bool

	//是否需要读取多段数据
	IsNeedReadMultiPart bool

	// 方法参数名称集合
	ParamNameList []string
}

// 检查IP是否合法
func (this *Handler) CheckIP(r *http.Request) *ResultStatus {
	if this.IsCheckIP && debugUtil.IsDebug() == false && ipMgr.IsIpValid(webUtil.GetRequestIP(r)) == false {
		return InvalidIP
	}

	return Success
}

// 检测参数
func (this *Handler) CheckParam(r *http.Request) *ResultStatus {
	// 在第一次使用参数时，调用ParseForm方法
	r.ParseForm()

	if this.IsNeedReadMultiPart {
		//上传文件，上传时用的类型是multipart/form-data，所以必须调用ParseMultipartForm才会读取流中的表单数据
		r.ParseMultipartForm(32 << 20) // 32 MB
	}

	for _, name := range this.ParamNameList {
		if r.PostForm[name] == nil || len(r.PostForm[name]) == 0 {
			return APIParamError
		}
	}

	return Success
}

/*
	创建新的请求方法对象
	参数：
		path:请求路径
		handlerFunc:处理方法
		isCheckIP:是否需要检查IP
		isNeedReadMultiPart:是否需要读取多段数据
		paramNameList:需要的参数名称集合
	返回：
		请求方法对象

*/
func NewMultiPartHandler(path string,
	handlerFunc func(http.ResponseWriter, *http.Request) *ResponseObject,
	isCheckIP bool,
	isNeedReadMultiPart bool,
	paramNameList ...string) *Handler {

	return &Handler{
		Path:                path,
		HandlerFunc:         handlerFunc,
		IsCheckIP:           isCheckIP,
		IsNeedReadMultiPart: isNeedReadMultiPart,
		ParamNameList:       paramNameList,
	}
}

/*
	创建新的请求方法对象
	参数：
		path:请求路径
		handlerFunc:处理方法
		isCheckIP:是否需要检查IP
		paramNameList:需要的参数名称集合
	返回：
		请求方法对象

*/
func NewHandler(path string,
	handlerFunc func(http.ResponseWriter, *http.Request) *ResponseObject,
	isCheckIP bool,
	paramNameList ...string) *Handler {

	return NewMultiPartHandler(path,
		handlerFunc,
		isCheckIP,
		false,
		paramNameList...)
}
