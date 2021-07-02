package calc

import (
	"github.com/znyh/class/qp/majong/calc/hulib"
)

var (
	checkSpecialCardType = make([]func(*CardsSet) bool, 0)
)

func init() {
	//若不用特殊牌型，注释下
	checkSpecialCardType = []func(*CardsSet) bool{Check7dui, Check13Yao}
}

type CardsSet struct {
	elem     []int          //原始的cardsIndex
	count    []int          //去掉赖子后的cardsIndex
	weaves   []TagWeaveItem //weave
	size     int
	guiIndex int
	guiNum   int
	ting     []int
	hu       bool
}

func NewCardsSet(cbCardsIndex []int, weaves []TagWeaveItem, guiIndex int) *CardsSet {
	set := &CardsSet{}

	set.guiIndex = guiIndex
	set.elem = make([]int, MAXINDEX)
	copy(set.elem, cbCardsIndex)

	set.weaves = make([]TagWeaveItem, len(weaves))
	copy(set.weaves, weaves)

	set.refresh()

	return set
}

func Check7dui(set *CardsSet) bool {
	return set.Check7dui()
}

func Check13Yao(set *CardsSet) bool {
	return set.Check13Yao()
}

func (set *CardsSet) refresh() {
	set.size = 0
	set.count = make([]int, MAXINDEX)
	set.ting = []int(nil)

	copy(set.count, set.elem)

	if set.guiIndex >= 0 && set.guiIndex < MAXINDEX {
		set.guiNum = set.count[set.guiIndex]
		set.count[set.guiIndex] = 0
	} else {
		set.guiIndex = MAXINDEX
	}

	for _, c := range set.elem {
		set.size += c
	}

	if set.size%3 == 2 {
		set.hu = set.checkHu()
	} else if set.size%3 == 1 {
		set.ting = set.checkTing()
	}
}

func (set *CardsSet) checkHu() bool {
	if hulib.CheckHu(set.elem[:], set.guiIndex) {
		return true
	}
	for _, v := range checkSpecialCardType {
		if ok := v(set); ok {
			return true
		}
	}
	return false
}

func (set *CardsSet) checkTing() []int {
	ting := []int(nil)

	count := make([]int, MAXINDEX)
	copy(count, set.elem)

	for i := 0; i < MAXINDEX; i++ {
		if count[i] >= 4 {
			continue
		}
		count[i]++
		if hulib.CheckHu(count, set.guiIndex) {
			ting = append(ting, i)
		} else {
			for _, v := range checkSpecialCardType {
				if ok := v(set); ok {
					ting = append(ting, i)
				}
			}
		}
		count[i]--
	}
	return ting
}

func (set *CardsSet) add(index int) {
	if index < MAXINDEX && index >= 0 {
		set.elem[index]++
		set.refresh()
	}
}

func (set *CardsSet) del(index int) {
	if index < MAXINDEX && index >= 0 {
		if set.elem[index] > 0 {
			set.elem[index]--
			set.refresh()
		}
	}
}

func (set *CardsSet) CalcHandsColor() map[int]int {
	colors := map[int]int{}
	for i, c := range set.elem {
		if c > 0 && i != set.guiIndex {
			colors[i/9] = 1
		}
	}
	return colors
}

func (set *CardsSet) CalcWeavesColor() map[int]int {
	colors := map[int]int{}
	for _, weave := range set.weaves {
		color := (weave.cbCenterCard & MASKCOLOR) >> 4
		colors[color] = 1
	}
	return colors
}

func (set *CardsSet) Check7dui() bool {
	sum := 0
	need := 0
	count, guiNum := set.calcGuiNum()

	for i := 0; i < MAXINDEX; i++ {
		c := count[i]
		sum += c
		if (c > 0) && (c%2 != 0) {
			need++
		}
	}
	return (sum+guiNum == 14) && (guiNum >= need)
}

func (set *CardsSet) CheckSuper7dui() (bool, int) {
	count, guiNum := set.calcGuiNum()
	left := guiNum
	sum, dui, need, super := 0, 0, 0, 0

	for i := 0; i < MAXINDEX; i++ {
		if c := count[i]; c != 0 {
			sum += c
			if c <= 2 {
				dui = dui + 1
				left = left - (2 - c)
				need = need + (2 - c)
			} else {
				super = super + 1
				left = left - (4 - c)
				need = need + (4 - c)
			}
		}
	}

	if left/2 >= dui && left >= 0 {
		super = super + dui + (left/2-dui)/2
	} else if left/2 < dui && left >= 0 {
		super = super + left/2
	}

	return (sum+guiNum == 14) && (guiNum >= need) && (super >= 0), super
}

func (set *CardsSet) Check13Yao() bool {
	//13yao
	SHISANYAO_CARDS := []int{0, 8, 9, 17, 18, 26, 27, 28, 29, 30, 31, 32, 33}

	count, guiNum := set.calcGuiNum()
	sum, need := 0, 0
	eye := false

	for _, i := range SHISANYAO_CARDS {
		if c := count[i]; c > 0 {
			if c > 2 {
				return false
			} else if c == 2 {
				if eye {
					return false
				} else {
					eye = true
				}
			}
			sum = sum + c
		} else {
			need++
		}
	}

	if eye {
		return (need == guiNum) && (sum+guiNum == 14)
	} else {
		return (need+1 == guiNum) && (sum+guiNum == 14)
	}
}

func (set *CardsSet) CheckPengPengHu() bool {
	for _, weave := range set.weaves {
		weaveType := weave.cbWeaveKind
		if weaveType == OpEatLeft || weaveType == OpEatCenter || weaveType == OpEatRight {
			return false
		}
	}

	count, guiNum := set.calcGuiNum()
	need := 0
	for i := 0; i < MAXINDEX; i++ {
		c := count[i]
		if c > 0 && c < 4 {
			need = need + 3 - c
		} else if c == 4 {
			need = need + 2
		}
	}
	if (guiNum-need-2 >= 0) || (guiNum-need+1 >= 0) {
		return (guiNum-need-2)%3 == 0 || (guiNum-need+1)%3 == 0
	}
	return false
}

func (set *CardsSet) CheckKaCard(KaIndex int) bool {
	if KaIndex < 0 || KaIndex >= MAXINDEX {
		return false
	}

	if set.elem[KaIndex] <= 0 {
		return false
	}

	flag, huNum := 0, 0
	count := make([]int, MAXINDEX)
	copy(count, set.elem)
	count[KaIndex]--

	for i := 0; i < MAXINDEX; i++ {
		count[i]++
		if hulib.CheckHu(count, set.guiIndex) {
			if i == set.guiIndex {
				flag = 1
			}
			huNum++
		}
		count[i]--
	}
	return huNum == (1 + flag)
}

func (set *CardsSet) CheckYiTiaoLong() bool {
	for color := 0; color < 3; color++ {
		need := 0
		guiIndex := set.guiIndex
		count, guiNum := set.calcGuiNum()

		for card := 0 + 9*color; card < 9+9*color; card++ {
			if count[card] == 0 {
				need++
			} else {
				count[card]--
			}
		}
		if guiNum-need >= 0 {
			if guiIndex < MAXINDEX && guiIndex >= 0 {
				count[guiIndex] = guiNum - need
			}
			if hulib.CheckHu(count, guiIndex) {
				return true
			}
		}
	}
	return false
}

func (set *CardsSet) CheckHuaLong() bool {
	if set.size < 9 {
		return false
	}
	if !hulib.CheckHu(set.elem, set.guiIndex) {
		return false
	}
	for c := 0; c < 3; c++ {
		for c1 := 0; c1 < 3; c1++ {
			if c != c1 {
				guiIndex := set.guiIndex
				tmpCards := make([]int, MAXINDEX)
				copy(tmpCards, set.elem)

				colors := [3]int{c, c1, 3 - c - c1}

				guiNum := 0
				if guiIndex < 34 && guiIndex >= 0 {
					guiNum = tmpCards[guiIndex]
				}

				need := set.removeAndGetNeed(tmpCards, colors)

				if need > guiNum {
					continue
				}

				if guiIndex < 34 && guiIndex >= 0 && guiNum >= need {
					tmpCards[guiIndex] -= need
				}

				if hulib.CheckHu(tmpCards, guiIndex) {
					return true
				}
			}
		}
	}
	return false
}

func (set *CardsSet) removeAndGetNeed(tmpCards []int, Colors [3]int) int {
	need := 0
	for i := 0; i < 9; i++ {
		color := i / 3
		index := Colors[color] * 9

		if tmpCards[index+i] <= 0 {
			need++
		} else {
			tmpCards[index+i]--
		}
	}
	return need
}

func (set *CardsSet) CheckQingYiSe() bool {
	colors := set.CalcHandsColor()
	for _, weave := range set.weaves {
		color := (weave.cbCenterCard & MASKCOLOR) >> 4
		colors[color] = 1
	}
	return (colors[0]+colors[1]+colors[2] == 1) && colors[3] == 0
}

func (set *CardsSet) CheckHunYiSe() bool {
	colors := set.CalcHandsColor()
	for _, weave := range set.weaves {
		color := (weave.cbCenterCard & MASKCOLOR) >> 4
		colors[color] = 1
	}
	return (colors[0]+colors[1]+colors[2] == 1) && colors[3] == 1
}

func (set *CardsSet) CheckZiYiSe() bool {
	colors := set.CalcHandsColor()
	for _, weave := range set.weaves {
		color := (weave.cbCenterCard & MASKCOLOR) >> 4
		colors[color] = 1
	}
	return (colors[0]+colors[1]+colors[2] == 0) && colors[3] == 1
}

func (set *CardsSet) calcGuiNum() ([]int, int) {
	count := make([]int, MAXINDEX)
	copy(count, set.count)
	return count, set.guiNum
}

func (set *CardsSet) CheckEat(currIndex int) int {
	if currIndex >= MAXINDEX && currIndex < 0 {
		return 0
	}

	cbExcursion := [3]int{0, 1, 2}
	cbItemKind := [3]int{OpEatLeft, OpEatCenter, OpEatRight}
	cbEatKind, cbFirstIndex := 0, 0

	for i := 0; i < 3; i++ {
		cbValueIndex := currIndex % 9
		if cbValueIndex >= cbExcursion[i] && cbValueIndex-cbExcursion[i] <= 6 {
			cbFirstIndex = currIndex - cbExcursion[i]
			if set.guiIndex >= cbFirstIndex && set.guiIndex <= cbFirstIndex+2 {
				continue
			}
			if currIndex != cbFirstIndex && set.elem[cbFirstIndex] == 0 {
				continue
			}
			if currIndex != cbFirstIndex+1 && set.elem[cbFirstIndex+1] == 0 {
				continue
			}
			if currIndex != cbFirstIndex+2 && set.elem[cbFirstIndex+2] == 0 {
				continue
			}
			cbEatKind |= cbItemKind[i]
		}
	}
	return cbEatKind
}
