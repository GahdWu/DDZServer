package rpcServer

import (
	"github.com/Gahd/DDZServer/src/model/player"
	. "github.com/Gahd/DDZServer/src/model/responseObject"
)

var (
	// 服务器监听地址
	serverAddress string

	// 获取玩家对象的方法
	getPlayerFunc func(*Client, string) (*player.Player, *ResultStatus)

	// 是否DEBUG模式
	debug bool
)

// 设置Config信息
// _serverAddress:服务器监听地址
// _getPlayerFunc:获取玩家对象的方法
// _debug:是否DEBUG模式
func SetConfig(_serverAddress string, _getPlayerFunc func(*Client, string) (*player.Player, *ResultStatus), _debug bool) {
	serverAddress = _serverAddress
	getPlayerFunc = _getPlayerFunc
	debug = _debug
}
