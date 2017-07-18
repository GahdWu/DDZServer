package responseObject

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
	ParamInValid        = newResultStatus(-5, "ParamInValid", "方法未定义")
	DataFormatError     = newResultStatus(-6, "DataFormatError", "数据格式错误")
	PlayerNotLogin      = newResultStatus(-7, "PlayerNotLogin", "玩家未登录")
	ParamIsEmpty        = newResultStatus(-8, "ParamIsEmpty", "参数为空")
	ParamNotMatch       = newResultStatus(-9, "ParamNotMatch", "参数不匹配")
	ParamTypeError      = newResultStatus(-10, "ParamTypeError", "参数类型错误")
	OnlySupportPOST     = newResultStatus(-11, "OnlySupportPOST", "只支持POST")
	APINotDefined       = newResultStatus(-12, "APINotDefined", "API未定义")
	APIDataError        = newResultStatus(-13, "APIDataError", "API数据错误")
	APIParamError       = newResultStatus(-14, "APIParamError", "API参数错误")
	InvalidIP           = newResultStatus(-15, "InvalidIP", "IP无效")
	ReloadError         = newResultStatus(-16, "ReloadError", "重新加载出错")
	TokenNotExist       = newResultStatus(-17, "TokenNotExist", "Token不存在")
	AppNotExist         = newResultStatus(-18, "AppNotExist", "应用程序不存在")
	AppNotOpen          = newResultStatus(-19, "AppNotOpen", "App未开启")
	ChannelNotExist     = newResultStatus(-20, "ChannelNotExist", "渠道配置不存在")
	ChannelNotOpen      = newResultStatus(-21, "ChannelNotOpen", "渠道未开启")
	UserNotExist        = newResultStatus(-22, "UserNotExist", "用户不存在")
	UserIsForbidden     = newResultStatus(-23, "UserIsForbidden", "用户被封号")
	ConfigError         = newResultStatus(-24, "ConfigError", "配置错误")
	VerifyFail          = newResultStatus(-25, "VerifyFail", "SDK验证失败")
	TokenInvalid        = newResultStatus(-26, "TokenInvalid", "Token无效")
	UserDataException   = newResultStatus(-27, "UserDataException", "用户数据异常")
	OrderNotExist       = newResultStatus(-28, "OrderNotExist", "订单不存在")
	RemoteServerError   = newResultStatus(-29, "RemoteServerError", "远程服务器出错")
	SignError           = newResultStatus(-30, "SignError", "签名错误")
	DBConnectFail       = newResultStatus(-31, "DBConnectFail", "数据库连接失败")
	DBConfigNotExists   = newResultStatus(-32, "DBConfigNotExist", "数据库配置不存在")
	UploadFailed        = newResultStatus(-33, "UploadFailed", "上传失败")
	PlayerAlreadyInRoom = newResultStatus(-34, "PlayerAlreadyInRoom", "玩家已在房间中")
	PlayerAlreadyInGame = newResultStatus(-35, "PlayerAlreadyInGame", "玩家已在游戏中")
	RoomNotExists       = newResultStatus(-36, "RoomNotExists", "房间不存在")
	RoomIsFull          = newResultStatus(-37, "RoomIsFull", "房间已满")
	PlayerNotInRoom     = newResultStatus(-38, "PlayerNotInRoom", "玩家已不在房间中")
	CanNotFindHall      = newResultStatus(-39, "CanNotFindHall", "找不到对应的大厅")
	PlayerOnReady       = newResultStatus(-40, "PlayerOnReady", "玩家已准备")
)
