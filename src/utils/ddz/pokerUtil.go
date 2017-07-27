package ddz

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	. "github.com/Gahd/DDZServer/src/model/ddz/common"
	. "github.com/Gahd/DDZServer/src/model/ddz/poker"
)

var (
	pokerMutex = sync.Mutex{}
	pokers     = []*Poker{}
)

func init() {
	pokerMutex.Lock()
	defer pokerMutex.Unlock()

	//所有数字
	pokerColors := []PokerColor{HeiTao, HongTao, MeiHua, FangPian}
	pokerNums := []PokerNum{Three, Four, Five, Six, Seven, Eight, Nine, Ten, J, Q, K, A, Two}

	for _, n := range pokerNums {
		for _, c := range pokerColors {
			pokers = append(pokers, NewPoker(c, n))
		}
	}

	//大小王
	pokers = append(pokers, NewPoker(NoneColor, SKing))
	pokers = append(pokers, NewPoker(NoneColor, BKing))

	//展示所有扑克
	//	ShowPokers(pokers)
}

func GetFullPokers() []*Poker {
	pokerMutex.Lock()
	defer pokerMutex.Unlock()

	result := make([]*Poker, len(pokers))
	copy(result, pokers)

	return result
}

func GetRandFullPokers() []*Poker {
	result := GetFullPokers()
	RandPokers(result)
	return result
}

func RandPokers(source []*Poker) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	pokerSize := len(source)

	for i := 0; i < pokerSize; i++ {
		randIndex1 := r.Intn(pokerSize)
		source[randIndex1], source[i] = source[i], source[randIndex1]
	}
}

func SliceDDZPokers(source []*Poker) (p1 []*Poker, p2 []*Poker, p3 []*Poker, threePoker []*Poker) {
	p1 = make([]*Poker, 17) //17张
	p2 = make([]*Poker, 17) //17张
	p3 = make([]*Poker, 17) //17张

	threePoker = make([]*Poker, 3) //地主牌

	if len(source) != 54 {
		panic(fmt.Sprintf("牌数量[%d]不足54", len(source)))
	}

	copy(p1, source[:17])
	copy(p2, source[17:34])
	copy(p3, source[34:51])
	copy(threePoker, source[51:])

	return
}

func ShowPokers(source []*Poker) {
	for _, item := range source {
		if item == nil {
			fmt.Println("nil")
			continue
		}
		fmt.Println(item.ToString())
	}
}

//获取牌类型
func GetPokerGroupTypeByPokers(pokers []*Poker) PokerGroupType {
	length := len(pokers)

	switch {
	case length == 1: //只有一张牌，为单牌
		return DanPai
	case length == 2:
		if pokers[0].GetNum() == SKing && pokers[1].GetNum() == BKing {
			return WangZha //排序后，依次为小王和大王，则为王炸
		} else if pokers[0].GetNum() == pokers[1].GetNum() {
			return DuiZi //两张牌，并且数字相同为对子
		}
		break
	case length == 3: //三张牌，并且数字相同为三张不带
		if pokers[0].GetNum() == pokers[1].GetNum() && pokers[0].GetNum() == pokers[2].GetNum() {
			return SanDai
		}
		break
	case length == 4: //四张牌，三带一、炸弹
		if pokers[0].GetNum() == pokers[3].GetNum() {
			return ZhaDan //排序后第一张和最后一张相同，炸弹
		}

		if isSanDaiYiOrDui(pokers, 1) {
			return SanDaiYi
		}
		break
	case length >= 5: //五张牌以上，顺子、三带对、连对、飞机
		if isShunZi(pokers) {
			return ShunZi
		}

		if length == 5 {
			if isSanDaiYiOrDui(pokers, 2) {
				return SanDaiDui
			}
		}

		if isLianDui(pokers) {
			return LianDui
		}

		if isFeiJi(pokers) {
			return FeiJi
		}

		break
	}

	return UnkownGroupType
}

func getPokersSameNumInfo(pokers []*Poker) map[PokerNum]int {
	info := make(map[PokerNum]int)

	for _, p := range pokers {
		if s, isExists := info[p.GetNum()]; !isExists {
			info[p.GetNum()] = 1
		} else {
			info[p.GetNum()] = s + 1
		}
	}

	return info
}

func getNumInfoNumArray(numInfo map[PokerNum]int) []PokerNum {
	nums := []PokerNum{}
	for pn, _ := range numInfo {
		nums = append(nums, pn)
	}

	return nums
}

func orderPokerNumsByAsc(nums []PokerNum) []PokerNum {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if int(nums[i]) > int(nums[j]) {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	return nums
}

//是否是三带一或三代二
func isSanDaiYiOrDui(pokers []*Poker, daiSize int) bool {
	sameNumInfo := getPokersSameNumInfo(pokers)
	if len(sameNumInfo) != 2 {
		return false
	}

	var (
		isHaveThree = false
		isHaveOne   = false
	)

	for _, s := range sameNumInfo {
		switch s {
		case daiSize:
			isHaveOne = true
			break
		case 3:
			isHaveThree = true
			break
		}
	}

	return isHaveThree && isHaveOne
}

//是否是顺子，排序后每个数字+1等于后面的数字，排除2和王
func isShunZi(pokers []*Poker) bool {
	if len(pokers) < 5 {
		return false
	}
	num := -1
	for _, p := range pokers {
		switch p.GetNum() {
		case Two:
		case SKing:
		case BKing:
			return false //排除大小王和2
		}

		n := int(p.GetNum())
		if num != -1 && num+1 != n {
			return false
		}
		num = n
	}

	return true
}

//是否是连对，每个数字有两个，排序后每个数字+1等于后面的数字，排除2和王
func isLianDui(pokers []*Poker) bool {
	sameNumInfo := getPokersSameNumInfo(pokers)
	if len(sameNumInfo) < 2 {
		return false
	}

	num := -1
	for _, pn := range orderPokerNumsByAsc(getNumInfoNumArray(sameNumInfo)) {
		c := sameNumInfo[pn]
		if c != 2 {
			return false //不是对子
		}
		switch pn {
		case Two:
		case SKing:
		case BKing:
			return false //排除大小王和2
		}

		n := int(pn)
		if num != -1 && num+1 != n {
			return false
		}
		num = n
	}

	return true
}

//是否是飞机
func isFeiJi(pokers []*Poker) bool {
	sameNumInfo := getPokersSameNumInfo(pokers)

	var (
		threeSize = 0 //三个的数量
		oneSize   = 0 //是否是带对
		towSize   = 0 //是否是带单
	)

	num := -1
	for _, pn := range orderPokerNumsByAsc(getNumInfoNumArray(sameNumInfo)) {
		s := sameNumInfo[pn]
		switch s {
		case 1:
			oneSize++
			break
		case 2:
			towSize++
			break
		case 3:
			switch pn {
			case Two:
			case SKing:
			case BKing:
				return false //排除大小王和2
			}
			n := int(pn)
			if num != -1 && num+1 != n {
				return false
			}
			num = n
			threeSize++
			break
		}
	}

	if oneSize > 0 && towSize > 0 {
		return false //带了对又带了单
	}

	//带的单牌数量、对子数量与三张数量相同或不带
	return (oneSize == threeSize || towSize == threeSize || towSize+oneSize == 0) && threeSize > 1
}

//快速排序（排序10000个随机整数，用时约0.9ms）
func QuickSortPoker(pokers []*Poker) {
	recursionSortPoker(pokers, 0, len(pokers)-1)
}

//递归排序
func recursionSortPoker(pokers []*Poker, left int, right int) {
	if left < right {
		pivot := partitionSortPoker(pokers, left, right)
		recursionSortPoker(pokers, left, pivot-1)
		recursionSortPoker(pokers, pivot+1, right)
	}
}

//分区排序
func partitionSortPoker(pokers []*Poker, left int, right int) int {
	for left < right {
		for sortComparePoker(pokers, left, right) {
			right--
		}
		if left < right {
			pokers[left], pokers[right] = pokers[right], pokers[left]
			left++
		}

		for sortComparePoker(pokers, left, right) {
			left++
		}
		if left < right {
			pokers[left], pokers[right] = pokers[right], pokers[left]
			right--
		}
	}

	return left
}

//排序比较
func sortComparePoker(pokers []*Poker, left int, right int) bool {
	if left >= right {
		return false //分区排序循环完成
	}

	leftPoker := pokers[left]
	rightPoker := pokers[right]

	leftNumValue := int(leftPoker.GetNum())
	rightNumValue := int(rightPoker.GetNum())

	if leftNumValue != rightNumValue {
		return leftNumValue < rightNumValue
	}

	leftColorValue := int(leftPoker.GetColor())
	rightColorValue := int(rightPoker.GetColor())

	return leftColorValue < rightColorValue
}
