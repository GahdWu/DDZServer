package pokerGroup

import (
	"fmt"
	"testing"

	. "github.com/Gahd/DDZServer/src/model/ddz/common"
	. "github.com/Gahd/DDZServer/src/model/ddz/poker"
)

func TestPokerGroupType(t *testing.T) {
	var (
		pg     *PokerGroup
		pokers []*Poker
	)

	//-------单牌
	pokers = []*Poker{}
	pokers = append(pokers, NewPoker(MeiHua, J))

	pg = NewPokerGroup(pokers)

	pg.ShowAllPokers()
	if pg.GetPokerGroupType() != DanPai {
		t.Fatalf("单牌,牌型识别错误=>%s", pg.GetPokerGroupType().ToString())
	}
	fmt.Println()

	//-------对子
	pokers = []*Poker{}
	pokers = append(pokers, NewPoker(MeiHua, J))
	pokers = append(pokers, NewPoker(FangPian, J))

	pg = NewPokerGroup(pokers)

	pg.ShowAllPokers()
	if pg.GetPokerGroupType() != DuiZi {
		t.Fatalf("对子,牌型识别错误=>%s", pg.GetPokerGroupType().ToString())
	}
	fmt.Println()

	//-------王炸
	pokers = []*Poker{}
	pokers = append(pokers, NewPoker(NoneColor, SKing))
	pokers = append(pokers, NewPoker(NoneColor, BKing))

	pg = NewPokerGroup(pokers)

	pg.ShowAllPokers()
	if pg.GetPokerGroupType() != WangZha {
		t.Fatalf("王炸,牌型识别错误=>%s", pg.GetPokerGroupType().ToString())
	}
	fmt.Println()

	//-------三不带
	pokers = []*Poker{}
	pokers = append(pokers, NewPoker(MeiHua, Three))
	pokers = append(pokers, NewPoker(FangPian, Three))
	pokers = append(pokers, NewPoker(HeiTao, Three))

	pg = NewPokerGroup(pokers)

	pg.ShowAllPokers()
	if pg.GetPokerGroupType() != SanDai {
		t.Fatalf("三不带,牌型识别错误=>%s", pg.GetPokerGroupType().ToString())
	}
	fmt.Println()

	//-------三带一
	pokers = []*Poker{}
	pokers = append(pokers, NewPoker(MeiHua, Three))
	pokers = append(pokers, NewPoker(FangPian, Three))
	pokers = append(pokers, NewPoker(HeiTao, Three))
	pokers = append(pokers, NewPoker(MeiHua, Five))

	pg = NewPokerGroup(pokers)
	pg.ShowAllPokers()
	if pg.GetPokerGroupType() != SanDaiYi {
		t.Fatalf("三带一,牌型识别错误=>%s", pg.GetPokerGroupType().ToString())
	}
	fmt.Println()

	//-------炸弹
	pokers = []*Poker{}
	pokers = append(pokers, NewPoker(MeiHua, Five))
	pokers = append(pokers, NewPoker(FangPian, Five))
	pokers = append(pokers, NewPoker(HeiTao, Five))
	pokers = append(pokers, NewPoker(HongTao, Five))

	pg = NewPokerGroup(pokers)
	pg.ShowAllPokers()
	if pg.GetPokerGroupType() != ZhaDan {
		t.Fatalf("炸弹,牌型识别错误=>%s", pg.GetPokerGroupType().ToString())
	}
	fmt.Println()

	//-------三带对
	pokers = []*Poker{}
	pokers = append(pokers, NewPoker(MeiHua, Five))
	pokers = append(pokers, NewPoker(FangPian, Five))
	pokers = append(pokers, NewPoker(HeiTao, Five))
	pokers = append(pokers, NewPoker(HongTao, Three))
	pokers = append(pokers, NewPoker(HeiTao, Three))

	pg = NewPokerGroup(pokers)
	pg.ShowAllPokers()
	if pg.GetPokerGroupType() != SanDaiDui {
		t.Fatalf("三带对,牌型识别错误=>%s", pg.GetPokerGroupType().ToString())
	}
	fmt.Println()

	//-------顺子
	pokers = []*Poker{}
	pokers = append(pokers, NewPoker(MeiHua, Five))
	pokers = append(pokers, NewPoker(FangPian, Six))
	pokers = append(pokers, NewPoker(HeiTao, Seven))
	pokers = append(pokers, NewPoker(HongTao, Eight))
	pokers = append(pokers, NewPoker(HeiTao, Nine))
	pokers = append(pokers, NewPoker(HeiTao, Ten))

	pg = NewPokerGroup(pokers)
	pg.ShowAllPokers()
	if pg.GetPokerGroupType() != ShunZi {
		t.Fatalf("顺子,牌型识别错误=>%s", pg.GetPokerGroupType().ToString())
	}
	fmt.Println()

	//-------连对
	pokers = []*Poker{}
	pokers = append(pokers, NewPoker(MeiHua, Five))
	pokers = append(pokers, NewPoker(HeiTao, Six))
	pokers = append(pokers, NewPoker(MeiHua, Seven))
	pokers = append(pokers, NewPoker(HongTao, Six))
	pokers = append(pokers, NewPoker(FangPian, Five))
	pokers = append(pokers, NewPoker(HeiTao, Seven))

	pg = NewPokerGroup(pokers)
	pg.ShowAllPokers()
	if pg.GetPokerGroupType() != LianDui {
		t.Fatalf("连对,牌型识别错误=>%s", pg.GetPokerGroupType().ToString())
	}
	fmt.Println()

	//-------飞机
	pokers = []*Poker{}
	pokers = append(pokers, NewPoker(FangPian, Seven))
	pokers = append(pokers, NewPoker(HeiTao, Six))
	pokers = append(pokers, NewPoker(MeiHua, Seven))
	pokers = append(pokers, NewPoker(HongTao, Six))
	pokers = append(pokers, NewPoker(FangPian, Six))
	pokers = append(pokers, NewPoker(HeiTao, Seven))
	pokers = append(pokers, NewPoker(HeiTao, Eight))
	pokers = append(pokers, NewPoker(FangPian, Eight))
	pokers = append(pokers, NewPoker(HongTao, Eight))

	pokers = append(pokers, NewPoker(FangPian, Five))
	pokers = append(pokers, NewPoker(HeiTao, J))
	pokers = append(pokers, NewPoker(HeiTao, A))

	pg = NewPokerGroup(pokers)
	pg.ShowAllPokers()
	if pg.GetPokerGroupType() != FeiJi {
		t.Fatalf("飞机,牌型识别错误=>%s", pg.GetPokerGroupType().ToString())
	}
	fmt.Println()
}
