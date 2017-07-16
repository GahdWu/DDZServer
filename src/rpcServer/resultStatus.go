package rpcServer

// 中心服务器响应结果状态对象
type ResultStatus struct {
	// 状态值(成功是0，非成功以负数来表示)
	Code int

	// 英文信息
	Message string

	// 中文描述
	Desc string `json:"-"`
}

func newResultStatus(code int, message, desc string) *ResultStatus {
	return &ResultStatus{
		Code: code,
		// 兼容客户端显示
		Message: desc,
		Desc:    desc,
	}
}

// 定义所有的响应结果的状态枚举值
var (
	Success             = newResultStatus(0, "Success", "成功")
	DataError           = newResultStatus(-1, "DataError", "数据错误")
	CacheError          = newResultStatus(-2, "CacheError", "缓存数据错误")
	DBError             = newResultStatus(-3, "DBError", "数据库错误")
	MethodNotDefined    = newResultStatus(-4, "MethodNotDefined", "方法未定义")
	NoTargetMethod      = newResultStatus(-5, "MethodNotDefined", "方法未定义")
	ParamIsEmpty        = newResultStatus(-6, "ParamIsEmpty", "参数为空")
	ParamNotMatch       = newResultStatus(-7, "ParamNotMatch", "参数不匹配")
	ParamInValid        = newResultStatus(-8, "ParamInValid", "输入参数无效")
	DataFormatError     = newResultStatus(-9, "DataFormatError", "请求数据格式错误")
	ParamTypeError      = newResultStatus(-10, "ParamTypeError", "参数类型错误")
	PlayerNotLogin      = newResultStatus(-11, "PlayerNotLogin", "玩家未登录")
	OnlySupportPOST     = newResultStatus(-12, "OnlySupportPOST", "只支持POST")
	APINotDefined       = newResultStatus(-13, "APINotDefined", "API未定义")
	APIDataError        = newResultStatus(-14, "APIDataError", "API数据错误")
	APIParamError       = newResultStatus(-15, "APIParamError", "API参数错误")
	InvalidIP           = newResultStatus(-16, "InvalidIP", "IP无效")
	ReloadError         = newResultStatus(-17, "ReloadError", "重新加载出错")
	TokenNotExist       = newResultStatus(-18, "TokenNotExist", "Token不存在")
	AppNotExist         = newResultStatus(-19, "AppNotExist", "应用程序不存在")
	AppNotOpen          = newResultStatus(-20, "AppNotOpen", "App未开启")
	ChannelNotExist     = newResultStatus(-21, "ChannelNotExist", "渠道配置不存在")
	ChannelNotOpen      = newResultStatus(-22, "ChannelNotOpen", "渠道未开启")
	UserNotExist        = newResultStatus(-23, "UserNotExist", "用户不存在")
	UserIsForbidden     = newResultStatus(-24, "UserIsForbidden", "用户被封号")
	ConfigError         = newResultStatus(-25, "ConfigError", "配置错误")
	VerifyFail          = newResultStatus(-26, "VerifyFail", "SDK验证失败")
	TokenInvalid        = newResultStatus(-27, "TokenInvalid", "Token无效")
	UserDataException   = newResultStatus(-28, "UserDataException", "用户数据异常")
	OrderNotExist       = newResultStatus(-29, "OrderNotExist", "订单不存在")
	RemoteServerError   = newResultStatus(-30, "RemoteServerError", "远程服务器出错")
	SignError           = newResultStatus(-31, "SignError", "签名错误")
	DBConnectFail       = newResultStatus(-32, "DBConnectFail", "数据库连接失败")
	DBConfigNotExists   = newResultStatus(-33, "DBConfigNotExist", "数据库配置不存在")
	PlayerAlreadyInRoom = newResultStatus(-34, "PlayerAlreadyInRoom", "玩家已在房间中")
	PlayerAlreadyInGame = newResultStatus(-35, "PlayerAlreadyInGame", "玩家已在游戏中")
)
