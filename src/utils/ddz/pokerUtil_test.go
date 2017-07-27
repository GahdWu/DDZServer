package ddz

import (
	"fmt"
	"testing"
)

func TestPokerUtil(t *testing.T) {

	//获取一副随机的牌
	fullPokers := GetRandFullPokers()

	//展示牌
	ShowPokers(fullPokers)

	//分牌
	p1, p2, p3, three := SliceDDZPokers(fullPokers)

	//排序
	QuickSortPoker(p1)
	QuickSortPoker(p2)
	QuickSortPoker(p3)
	QuickSortPoker(three)

	fmt.Printf("\n玩家1：%d\n", len(p1))
	if len(p1) != 17 {
		t.Fatalf("玩家1的牌不够")
	}
	ShowPokers(p1)

	fmt.Printf("\n玩家2：%d\n", len(p2))
	if len(p2) != 17 {
		t.Fatalf("玩家2的牌不够")
	}
	ShowPokers(p2)

	fmt.Printf("\n玩家3：%d\n", len(p3))
	if len(p3) != 17 {
		t.Fatalf("玩家3的牌不够")
	}
	ShowPokers(p3)

	fmt.Printf("\n地主牌：%d\n", len(three))
	if len(three) != 3 {
		t.Fatalf("地主牌不够")
	}
	ShowPokers(three)

}
