package player

import (
	"sync"

	. "github.com/Gahd/DDZServer/src/model/player"
)

var (
	players map[string]*Player = make(map[string]*Player)
	mutex   sync.Mutex
)

func GetPlayer(playerId string) *Player {

	player := func() *Player {
		mutex.Lock()
		defer mutex.Unlock()

		player, isExists := players[playerId]
		if !isExists {
			return nil
		}
		return player
	}()

	if player == nil {
		player = CreatePlayer(playerId)
	}

	return player
}

func CreatePlayer(playerId string) *Player {
	mutex.Lock()
	defer mutex.Unlock()
	player := NewPlayer(playerId, 1)

	players[playerId] = player

	return player
}
