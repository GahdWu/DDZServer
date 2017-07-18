package ddz

import (
	"net/http"

	"github.com/Gahd/DDZServer/src/bll/player"
	"github.com/Gahd/DDZServer/src/model/ddz"
	. "github.com/Gahd/DDZServer/src/model/player"
	web "github.com/Gahd/DDZServer/src/model/responseObject"
	"github.com/Gahd/DDZServer/src/webServer"
)

func init() {
	//注册
	webServer.RegisteHandler(webServer.NewHandler("/DDZ/GetRooms", getRooms, false,
		"PlayerId"))
	webServer.RegisteHandler(webServer.NewHandler("/DDZ/CreateRoom", createRoom, false,
		"PlayerId"))
	webServer.RegisteHandler(webServer.NewHandler("/DDZ/EnterRoom", enterRoom, false,
		"PlayerId", "RoomId"))
	webServer.RegisteHandler(webServer.NewHandler("/DDZ/ExitRoom", exitRoom, false,
		"PlayerId"))
	webServer.RegisteHandler(webServer.NewHandler("/DDZ/QuickStart", quickStart, false,
		"PlayerId"))
}

/*
获取房间
http://192.168.1.109:9999/DDZ/GetRooms
PlayerId=testPlayerA
*/
func getRooms(w http.ResponseWriter, r *http.Request) *web.ResponseObject {
	result := web.NewResponseObject()

	playerId := r.FormValue("PlayerId")
	//	sign := r.FormValue("Sign")

	//该接口为普通斗地主
	hallType := ddz.DDZ_Normal
	player := player.GetPlayer(playerId)

	switch {
	case player.GetPlayerStatus() == InGame:
		return result.SetResultStatus(web.PlayerAlreadyInGame)
	case player.GetPlayerStatus() == InRoomUnReady || player.GetRoomId() != "":
		return result.SetResultStatus(web.PlayerAlreadyInRoom)
	}

	hallManager := GetHallManager()
	hall := hallManager.GetHall(hallType)
	if hall == nil {
		return result.SetResultStatus(web.CanNotFindHall)
	}

	result.SetData(hall.AssembleToClient())

	return result
}

/*
创建房间
http://192.168.1.109:9999/DDZ/CreateRoom
PlayerId=testPlayerA
*/
func createRoom(w http.ResponseWriter, r *http.Request) *web.ResponseObject {
	result := web.NewResponseObject()

	playerId := r.FormValue("PlayerId")
	//	sign := r.FormValue("Sign")

	//该接口为普通斗地主
	hallType := ddz.DDZ_Normal
	player := player.GetPlayer(playerId)

	switch {
	case player.GetPlayerStatus() == InGame:
		return result.SetResultStatus(web.PlayerAlreadyInGame)
	case player.GetPlayerStatus() == InRoomUnReady || player.GetRoomId() != "":
		return result.SetResultStatus(web.PlayerAlreadyInRoom)
	}

	hallManager := GetHallManager()
	hall := hallManager.GetHall(hallType)
	if hall == nil {
		return result.SetResultStatus(web.CanNotFindHall)
	}

	room := hall.CreateRoom(player)
	result.SetData(room.AssembleToClientDetails())

	return result
}

/*
进入房间
http://192.168.1.109:9999/DDZ/EnterRoom
PlayerId=testPlayerA&RoomId=
*/
func enterRoom(w http.ResponseWriter, r *http.Request) *web.ResponseObject {
	result := web.NewResponseObject()

	playerId := r.FormValue("PlayerId")
	roomId := r.FormValue("RoomId")
	//	sign := r.FormValue("Sign")

	//该接口为普通斗地主
	hallType := ddz.DDZ_Normal
	player := player.GetPlayer(playerId)

	switch {
	case player.GetPlayerStatus() == InGame:
		return result.SetResultStatus(web.PlayerAlreadyInGame)
	case player.GetPlayerStatus() == InRoomUnReady || player.GetRoomId() != "":
		return result.SetResultStatus(web.PlayerAlreadyInRoom)
	}

	hallManager := GetHallManager()
	hall := hallManager.GetHall(hallType)
	if hall == nil {
		return result.SetResultStatus(web.CanNotFindHall)
	}

	room, status := hall.EnterRoom(player, roomId)
	if status != web.Success {
		return result.SetResultStatus(status)
	}

	result.SetData(room.AssembleToClientDetails())

	return result
}

/*
退出房间
http://192.168.1.109:9999/DDZ/ExitRoom
PlayerId=testPlayerA
*/
func exitRoom(w http.ResponseWriter, r *http.Request) *web.ResponseObject {
	result := web.NewResponseObject()

	playerId := r.FormValue("PlayerId")
	//	sign := r.FormValue("Sign")

	//该接口为普通斗地主
	hallType := ddz.DDZ_Normal
	player := player.GetPlayer(playerId)

	if player.GetRoomId() == "" {
		return result.SetResultStatus(web.PlayerNotInRoom)
	}

	switch player.GetPlayerStatus() {
	case InHall:
		return result.SetResultStatus(web.PlayerNotInRoom)
	}

	hallManager := GetHallManager()
	hall := hallManager.GetHall(hallType)
	if hall == nil {
		return result.SetResultStatus(web.CanNotFindHall)
	}

	hall.ExitRoom(player)

	return result
}

/*
快速进入
http://192.168.1.109:9999/DDZ/QuickStart
PlayerId=testPlayerA
*/
func quickStart(w http.ResponseWriter, r *http.Request) *web.ResponseObject {
	result := web.NewResponseObject()

	playerId := r.FormValue("PlayerId")

	//该接口为普通斗地主
	hallType := ddz.DDZ_Normal
	player := player.GetPlayer(playerId)

	switch player.GetPlayerStatus() {
	case InRoomReady:
	case InRoomUnReady:
		return result.SetResultStatus(web.PlayerAlreadyInRoom)
	case InGame:
		return result.SetResultStatus(web.PlayerAlreadyInGame)
	}

	hallManager := GetHallManager()
	hall := hallManager.GetHall(hallType)
	if hall == nil {
		return result.SetResultStatus(web.CanNotFindHall)
	}

	//随机一个未满的房间或创建一个新的房间
	room := hall.RandomNotFullRoom(player)

	//进入房间
	room.EnterRoom(player)

	result.SetData(room.AssembleToClientDetails())

	return result
}
