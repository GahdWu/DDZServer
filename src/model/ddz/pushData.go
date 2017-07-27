package ddz

import (
	"github.com/Gahd/DDZServer/src/config"
	"github.com/Gahd/DDZServer/src/model/common"
	. "github.com/Gahd/DDZServer/src/model/player"
	. "github.com/Gahd/DDZServer/src/model/responseObject"
	"github.com/Gahd/DDZServer/src/rpcServer"
	"moqikaka.com/goutil/logUtil"
)

// 向客户端推送玩家移动信息
func PushStatusInfo(player *Player, pushPlayers []*Player) {
	responseObj := NewResponseObject()

	// 组装数据
	data := make(map[string]interface{})
	data[common.PlayerId] = player.GetId()
	data[common.PlayerStatus] = player.GetPlayerStatus()

	responseObj.SetData(data)

	// 向每一个成员发送消息
	pushDataToClients(pushPlayers, responseObj, rpcServer.Con_HighPriority)
}

// 推送数据给客户端
// memberList：战场玩家对象列表
// responseObj：响应对象
// priority：消息的优先级
func pushDataToClients(pushPlayers []*Player, responseObj *ResponseObject, priority rpcServer.Priority) {
	// 向每一个成员发送消息
	for _, player := range pushPlayers {
		pushDataToClient(player, responseObj, priority)
	}
}

// 推送消息给客户端
// member：成员对象
// responseObj：响应对象
// priority：消息的优先级
func pushDataToClient(player *Player, responseObj *ResponseObject, priority rpcServer.Priority) {
	pushDataToPlayer(player.GetId(), player.GetName(), responseObj, priority)
}

// 向玩家推送数据
// playerId：玩家Id
// playerName：玩家名称
// responseObj：响应对象
// priority：消息的优先级
func pushDataToPlayer(playerId, playerName string, responseObj *ResponseObject, priority rpcServer.Priority) {
	if clientObj, ok := rpcServer.GetClientByPlayerId(playerId); ok {
		if config.GetBaseConfig().DEBUG {
			logUtil.DebugLog("玩家:%v,主动推送消息:%v", playerName, responseObj)
		}
		rpcServer.ResponseResult(clientObj, nil, responseObj, priority)
	}
}
