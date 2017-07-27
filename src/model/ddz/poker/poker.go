package poker

import (
	"fmt"

	. "github.com/Gahd/DDZServer/src/model/ddz/common"
)

//扑克
type Poker struct {
	color PokerColor
	num   PokerNum
}

func NewEmptyPoker() *Poker {
	return &Poker{}
}

func NewPoker(_color PokerColor, _num PokerNum) *Poker {
	return &Poker{
		color: _color,
		num:   _num,
	}
}

func (this *Poker) GetNum() PokerNum {
	return this.num
}

func (this *Poker) GetColor() PokerColor {
	return this.color
}

func (this *Poker) ToString() string {
	if this.color == NoneColor {
		return fmt.Sprintf("[%s%s]", this.color.ToString(), this.num.ToString())
	} else {
		return fmt.Sprintf("[%s %s]", this.color.ToString(), this.num.ToString())
	}
}
