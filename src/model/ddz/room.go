package ddz

import (
	"sync"

	"github.com/Gahd/DDZServer/src/model/common"
	. "github.com/Gahd/DDZServer/src/model/player"
	web "github.com/Gahd/DDZServer/src/model/responseObject"
)

type Room struct {
	id      string
	players map[string]*Player

	hallType HallType

	// 玩家锁
	mutex sync.Mutex

	closeFun func(roomId string) bool
}

func NewEmptyRoom() *Room {
	return &Room{}
}

//创建房间
func NewRoom(_id string, _hallType HallType, _closeFun func(roomId string) bool) *Room {
	return &Room{
		id:       _id,
		hallType: _hallType,
		players:  map[string]*Player{},
		closeFun: _closeFun,
	}
}

//进入房间
func (this *Room) EnterRoom(player *Player) *web.ResultStatus {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if this.IsFull() {
		return web.RoomIsFull //已满
	}

	//放入房间
	this.players[player.GetId()] = player

	//设置玩家的房间ID
	player.SetRoomId(this.id)

	//设置玩家状态
	player.SetPlayerStatus(InRoomUnReady)

	return web.Success
}

//退出房间
func (this *Room) ExitRoom(player *Player) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	//从房间中删除房间
	delete(this.players, player.GetId())
	player.SetRoomId("")

	//设置玩家状态
	player.SetPlayerStatus(InHall)

	//清理空房间
	if this.IsEmpty() {
		this.closeFun(this.id)
	}
}

//房间是否已满
func (this *Room) IsFull() bool {
	return this.players != nil && len(this.players) >= common.Con_MaxRoomPlayerCount
}

//是否是空房间
func (this *Room) IsEmpty() bool {
	return this.players == nil || len(this.players) == 0
}

func (this *Room) GetId() string {
	return this.id
}

func (this *Room) AssembleToClientDetails() interface{} {

	playersInfo := make([]interface{}, 0)

	for _, player := range this.players {
		playersInfo = append(playersInfo, player.AssembleToClient())
	}

	info := make(map[string]interface{})
	info[common.RoomId] = this.id
	info[common.RoomPlayers] = playersInfo

	return info
}

func (this *Room) AssembleToClient() interface{} {

	info := make(map[string]interface{})
	info[common.RoomId] = this.id
	info[common.PlayerCount] = len(this.players)

	return info
}
