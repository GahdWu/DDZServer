package poker

import (
	"fmt"
)

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
	quickSortPoker(_pokers)
	return &PokerGroup{
		pokers:         _pokers,
		pokerGroupType: GetPokerGroupTypeByPokers(_pokers),
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
	for _, item := range this.pokers {
		fmt.Println(item.ToString())
	}
	fmt.Println(this.pokerGroupType.ToString())
}

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

//比较
func (this *PokerGroup) Compare(target PokerGroup) bool {
	return false
}

//获取牌类型
func GetPokerGroupTypeByPokers(pokers []*Poker) PokerGroupType {
	length := len(pokers)

	switch {
	case length == 1: //只有一张牌，为单牌
		return DanPai
	case length == 2:
		if pokers[0].num == SKing && pokers[1].num == BKing {
			return ZhaDan //排序后，依次为小王和大王，则为王炸
		} else if pokers[0].num == pokers[1].num {
			return DuiZi //两张牌，并且数字相同为对子
		}
		break
	case length == 3: //三张牌，并且数字相同为三张不带
		if pokers[0].num == pokers[1].num && pokers[0].num == pokers[2].num {
			return SanDai
		}
		break
	case length == 4: //四张牌，三带一、炸弹
		if pokers[0].num == pokers[3].num {
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

//TODO:顺序有问题
func getPokersSameNumInfo(pokers []*Poker) map[PokerNum]int {
	info := make(map[PokerNum]int)

	for _, p := range pokers {
		if s, isExists := info[p.num]; !isExists {
			info[p.num] = 1
		} else {
			info[p.num] = s + 1
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

func sortPokers(pokers []*Poker) {
	for i := 0; i < len(pokers)-1; i++ {
		if int(pokers[i].num) > int(pokers[i+1].num) {
			pokers[i], pokers[i+1] = pokers[i+1], pokers[i]
		}
	}
}

//快速排序（排序10000个随机整数，用时约0.9ms）
func quickSortPoker(pokers []*Poker) {
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

	leftNumValue := int(leftPoker.num)
	rightNumValue := int(rightPoker.num)

	if leftNumValue != rightNumValue {
		return leftNumValue < rightNumValue
	}

	leftColorValue := int(leftPoker.color)
	rightColorValue := int(rightPoker.color)

	return leftColorValue < rightColorValue
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
		switch p.num {
		case Two:
		case SKing:
		case BKing:
			return false //排除大小王和2
		}

		n := int(p.num)
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
