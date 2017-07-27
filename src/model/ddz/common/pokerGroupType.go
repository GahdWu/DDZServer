package common

type PokerGroupType int

const (
	UnkownGroupType PokerGroupType = iota
	DanPai                         //单牌
	DuiZi                          //对子
	SanDai                         //三张不带
	SanDaiYi                       //三张带一张
	SanDaiDui                      //三张带对子
	FeiJi                          //飞机
	ShunZi                         //顺子
	LianDui                        //连对
	ZhaDan                         //炸弹
	WangZha                        //王炸
)

func (this PokerGroupType) ToString() string {
	switch this {
	case DanPai:
		return "单牌"
	case DuiZi:
		return "对子"
	case SanDai:
		return "三不带"
	case SanDaiYi:
		return "三带一"
	case SanDaiDui:
		return "三带对"
	case FeiJi:
		return "飞机"
	case ShunZi:
		return "顺子"
	case LianDui:
		return "连对"
	case ZhaDan:
		return "炸弹"
	}

	return "未知牌型"
}
