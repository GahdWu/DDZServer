package pokerGroup

import (
	"fmt"

	. "github.com/Gahd/DDZServer/src/model/ddz/common"
	. "github.com/Gahd/DDZServer/src/model/ddz/poker"
	util "github.com/Gahd/DDZServer/src/utils/ddz"
)

type PokerGroup struct {
	pokers         []*Poker
	pokerGroupType PokerGroupType
}

func NewEmptyPokerGroup() *PokerGroup {
	return &PokerGroup{}
}

func NewPokerGroup(_pokers []*Poker) *PokerGroup {
	//	sortPokers(_pokers)
	util.QuickSortPoker(_pokers)
	return &PokerGroup{
		pokers:         _pokers,
		pokerGroupType: util.GetPokerGroupTypeByPokers(_pokers),
	}
}

func (this *PokerGroup) GetPokers() []*Poker {
	return this.pokers
}

func (this *PokerGroup) GetPokerGroupType() PokerGroupType {
	return this.pokerGroupType
}

//显示所有扑克
func (this *PokerGroup) ShowAllPokers() {
	util.ShowPokers(this.pokers)
	fmt.Println(this.pokerGroupType.ToString())
}

//比较
func (this *PokerGroup) Compare(target PokerGroup) bool {
	return false
}
