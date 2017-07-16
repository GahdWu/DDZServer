package ddz

import (
	"encoding/json"
	"sync"

	"moqikaka.com/goutil/stringUtil"

	"github.com/Gahd/DDZServer/src/model/common"
	. "github.com/Gahd/DDZServer/src/model/player"
	web "github.com/Gahd/DDZServer/src/model/responseObject"
	"moqikaka.com/goutil/debugUtil"
	"moqikaka.com/goutil/mathUtil"
)

//大厅类型
type HallType int

const (
	DDZ_Normal HallType = iota //斗地主
)

type Hall struct {
	hallType HallType
	rooms    map[string]*Room

	//大厅锁
	mutex sync.Mutex
}

func NewEmptyHall() *Hall {
	return &Hall{}
}

func NewHall(_hallType HallType) *Hall {
	return &Hall{
		hallType: _hallType,
		rooms:    make(map[string]*Room),
	}
}

//获取所有房间
func (this *Hall) GetRooms() []*Room {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	rooms := make([]*Room, 0)

	for _, v := range this.rooms {
		rooms = append(rooms, v)
	}

	return rooms
}

//随机一个未满的房间
func (this *Hall) RandomNotFullRoom(player *Player) *Room {
	//随机房间的处理
	randRoom := func() *Room {
		this.mutex.Lock()
		defer this.mutex.Unlock()

		notFullRoomKeys := make([]string, 0)

		for k, v := range this.rooms {
			if !v.IsFull() {
				notFullRoomKeys = append(notFullRoomKeys, k)
			}
		}

		notFullRoomCount := len(notFullRoomKeys)
		if notFullRoomCount == 0 {
			return nil
		}

		randRoomKeyIndex := mathUtil.GetRandInt(notFullRoomCount)

		return this.rooms[notFullRoomKeys[randRoomKeyIndex]]
	}()

	if randRoom == nil {
		//没有未满的房间，则创建一个房间
		randRoom = this.CreateRoom(player)
	} else {
		this.ShowHallInfo()
	}

	return randRoom
}

func (this *Hall) CreateRoom(player *Player) *Room {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	var room = NewRoom(stringUtil.GetNewGUID(), this.hallType, this.CloseRoom)

	//创建房间则进入房间
	room.EnterRoom(player)

	//立即准备
	player.SetPlayerStatus(InRoomReady)

	//保存房间
	this.rooms[room.GetId()] = room

	this.ShowHallInfo()

	return room
}

func (this *Hall) EnterRoom(player *Player, roomId string) (*Room, *web.ResultStatus) {
	//玩家还在房间中，则先退出房间
	if player.GetRoomId() != "" {
		this.ExitRoom(player)
	}

	this.mutex.Lock()
	defer this.mutex.Unlock()
	room, isExists := this.rooms[roomId]
	if !isExists {
		return nil, web.RoomNotExists
	}

	status := room.EnterRoom(player)

	this.ShowHallInfo()

	return room, status
}

func (this *Hall) ExitRoom(player *Player) {
	if player.GetRoomId() == "" {
		return
	}

	var room *Room
	var isExists bool
	func() {
		this.mutex.Lock()
		defer this.mutex.Unlock()

		room, isExists = this.rooms[player.GetRoomId()]
		if !isExists {
			player.SetRoomId("")
			return
		}
	}() //避免退出房间时关闭房间，导致死锁

	if room != nil {
		room.ExitRoom(player)
	}

	this.ShowHallInfo()
}

func (this *Hall) CloseRoom(roomId string) bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	room, isExists := this.rooms[roomId]
	if !isExists {
		return false
	}

	if !room.IsEmpty() {
		return false
	}

	delete(this.rooms, roomId)

	this.ShowHallInfo()

	return true
}

func (this *Hall) AssembleToClient() interface{} {

	roomInfo := make([]interface{}, 0)
	for _, v := range this.rooms {
		roomInfo = append(roomInfo, v.AssembleToClient())
	}

	info := make(map[string]interface{})
	info[common.HallType] = this.hallType
	info[common.RoomCount] = len(this.rooms)
	info[common.RoomPlayerMaxCount] = common.Con_MaxRoomPlayerCount
	info[common.HallRooms] = roomInfo

	return info
}

func (this *Hall) AssembleToClientDetails() interface{} {

	roomInfo := make([]interface{}, 0)
	for _, v := range this.rooms {
		roomInfo = append(roomInfo, v.AssembleToClientDetails())
	}

	info := make(map[string]interface{})
	info[common.HallType] = this.hallType
	info[common.RoomCount] = len(this.rooms)
	info[common.RoomPlayerMaxCount] = common.Con_MaxRoomPlayerCount
	info[common.HallRooms] = roomInfo

	return info
}

func (this *Hall) ShowHallInfo() {
	infos := this.AssembleToClientDetails()
	buffer, err := json.Marshal(infos)
	if err != nil {
		debugUtil.Println(err.Error())
		return
	}

	debugUtil.Println(string(buffer))
}
