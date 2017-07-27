package poker

import (
	"fmt"
)

//扑克牌颜色
type PokerColor int

const (
	NoneColor PokerColor = iota //无颜色(大小王)
	FangPian                    //方片
	MeiHua                      //梅花
	HongTao                     //红桃
	HeiTao                      //黑桃
)

//扑克牌面数字
type PokerNum int

const (
	Three PokerNum = iota //3
	Four                  //4
	Five                  //5
	Six                   //6
	Seven                 //7
	Eight                 //8
	Nine                  //9
	Ten                   //10
	J                     //J
	Q                     //Q
	K                     //K
	A                     //J
	Two                   //2
	SKing                 //小王
	BKing                 //大王
)

func (this *PokerColor) ToString() string {
	switch *this {
	case NoneColor:
		return ""
	case HeiTao:
		return "黑桃"
	case HongTao:
		return "红桃"
	case MeiHua:
		return "梅花"
	case FangPian:
		return "方片"
	}

	panic("未知的花色")
}

func (this *PokerNum) ToString() string {
	switch *this {
	case Two:
		return "2"
	case Three:
		return "3"
	case Four:
		return "4"
	case Five:
		return "5"
	case Six:
		return "6"
	case Seven:
		return "7"
	case Eight:
		return "8"
	case Nine:
		return "9"
	case Ten:
		return "10"
	case J:
		return "J"
	case Q:
		return "Q"
	case K:
		return "K"
	case A:
		return "A"
	case SKing:
		return "小王"
	case BKing:
		return "大王"
	}

	panic("未知的牌")
}

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

func (this *Poker) ToString() string {
	if this.color == NoneColor {
		return fmt.Sprintf("[%s%s]", this.color.ToString(), this.num.ToString())
	} else {
		return fmt.Sprintf("[%s %s]", this.color.ToString(), this.num.ToString())
	}
}
