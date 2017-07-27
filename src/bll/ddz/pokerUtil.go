package ddz

import (
	"math/rand"
	"sync"
	"time"

	p "github.com/Gahd/DDZServer/src/model/ddz/poker"
)

var (
	pokerMutex = sync.Mutex{}
	pokers     = []*p.Poker{}
)

func init() {
	pokerMutex.Lock()
	defer pokerMutex.Unlock()

	//所有数字
	pokerColors := []p.PokerColor{p.HeiTao, p.HongTao, p.MeiHua, p.FangPian}
	pokerNums := []p.PokerNum{p.Three, p.Four, p.Five, p.Six, p.Seven, p.Eight, p.Nine, p.Ten, p.J, p.Q, p.K, p.A, p.Two}

	for _, n := range pokerNums {
		for _, c := range pokerColors {
			pokers = append(pokers, p.NewPoker(c, n))
		}
	}

	//大小王
	pokers = append(pokers, p.NewPoker(p.NoneColor, p.SKing))
	pokers = append(pokers, p.NewPoker(p.NoneColor, p.BKing))

	//展示所有扑克
	p.NewPokerGroup(pokers).ShowAllPokers()
}

func GetFullPokers() []*p.Poker {
	pokerMutex.Lock()
	defer pokerMutex.Unlock()

	result := make([]*p.Poker, len(pokers))
	copy(result, pokers)

	return result
}

func GetRandFullPokers() []*p.Poker {
	result := GetFullPokers()
	RandPokers(result)
	return result
}

func RandPokers(source []*p.Poker) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	pokerSize := len(source)

	for i := 0; i < pokerSize; i++ {
		randIndex1 := r.Intn(pokerSize)
		source[randIndex1], source[i] = source[i], source[randIndex1]
	}
}
