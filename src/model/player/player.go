package player

import (
	"github.com/Gahd/DDZServer/src/model/common"
)

//玩家状态
type PlayerStatus int

const (
	InHall        PlayerStatus = iota //大厅中
	InRoomUnReady                     //房间中,未准备
	InRoomReady                       //房间中,已准备
	InGame                            //房间中,游戏中
)

// 玩家对象
type Player struct {
	// 玩家Id
	id string

	// 玩家对应的客户端Id
	clientId int32

	//玩家状态
	playerStatus PlayerStatus

	//玩家当前的房间ID
	roomId string
}

func (this *Player) GetId() string {
	return this.id
}

func (this *Player) GetRoomId() string {
	return this.roomId
}

func (this *Player) GetClientId() int32 {
	return this.clientId
}

func (this *Player) GetPlayerStatus() PlayerStatus {
	return this.playerStatus
}

func (this *Player) SetPlayerStatus(status PlayerStatus) {
	this.playerStatus = status
}

func (this *Player) SetRoomId(roomId string) {
	this.roomId = roomId
}

func NewPlayer(_id string, _clientId int32) *Player {
	return &Player{
		id:           _id,
		clientId:     _clientId,
		playerStatus: InHall,
		roomId:       "",
	}
}

func (this *Player) AssembleToClient() interface{} {
	info := make(map[string]interface{})
	info[common.PlayerId] = this.id
	info[common.PlayerStatus] = this.playerStatus

	return info
}
