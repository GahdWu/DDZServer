package ddz

import (
	"github.com/Gahd/DDZServer/src/model/player"
	. "github.com/Gahd/DDZServer/src/model/responseObject"
	"github.com/Gahd/DDZServer/src/rpcServer"
)

type DDZModel int

func init() {
	rpcServer.RegisterFunction(new(DDZModel))
}

// 准备
// playerObj:玩家对象
func (this *DDZModel) Ready(playerObj *player.Player) *ResponseObject {
	result := NewResponseObject()

	switch {
	case playerObj.GetPlayerStatus() == player.InRoomReady: //是否在房间中且已准备
		return result.SetResultStatus(PlayerOnReady)
	case playerObj.GetRoomId() == "" ||
		playerObj.GetPlayerStatus() != player.InRoomUnReady ||
		playerObj.GetPlayerStatus() == player.InHall:
		return result.SetResultStatus(PlayerNotInRoom) //不在房间中
	case playerObj.GetPlayerStatus() == player.InGame: //是否已开始游戏
		return result.SetResultStatus(PlayerAlreadyInGame)
	}

	//准备
	playerObj.SetPlayerStatus(player.InRoomReady)

	return result
}
