package common

//扑克牌颜色
type PokerColor int

const (
	NoneColor PokerColor = iota //无颜色(大小王)
	FangPian                    //方片
	MeiHua                      //梅花
	HongTao                     //红桃
	HeiTao                      //黑桃
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
