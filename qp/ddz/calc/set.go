package calc

import (
	"sort"

	"github.com/znyh/class/qp/base"
)

/*
	一组牌定义 0-20张 card等级范围：[1~16)
*/

// CardsSet 一组牌
type CardsSet struct {
	seqs   []int32 // 不会排序
	cards  []Card
	count  [16]int
	cnt    map[int]int
	colors map[int]int
	ct     int
	length int
	level  int
	calced bool
}

func NewCardsSet(seqs []int32) *CardsSet {
	set := &CardsSet{}

	set.seqs = make([]int32, len(seqs))
	copy(set.seqs, seqs)
	set.refresh()

	return set
}

func (set *CardsSet) Type() int {
	set.calc()
	return set.ct
}

func (set *CardsSet) GetCountMap() [16]int {
	return set.count
}

func (set *CardsSet) refresh() {
	set.calced = false
	set.cards = make([]Card, len(set.seqs))
	set.count = [16]int{}
	set.cnt = map[int]int{}
	set.colors = map[int]int{}

	for i := 0; i < len(set.seqs); i++ {
		set.cards[i].Seq = set.seqs[i]
		set.cards[i].calc()
		set.count[set.cards[i].level]++
		set.colors[set.cards[i].color] = 1
	}

	for _, c := range set.count {
		if c > 0 {
			set.cnt[c]++
		}
	}

	sort.Slice(set.cards, func(i, j int) bool {
		return set.cards[i].level > set.cards[j].level
	})

	set.calc()
}

func (set *CardsSet) calc() {
	if set.calced {
		return
	}
	set.calced = true

	for _, v := range checkCardType[len(set.cards)] {
		if v(set) {
			break
		}
	}
}

func (set *CardsSet) IsBigger(other *CardsSet) bool {
	set.calc()
	other.calc()

	if set.ct == CtUnknown ||
		other.ct == CtUnknown {
		return false
	}

	if set.ct == other.ct {
		if set.length == other.length && set.level > other.level {
			return true
		}
		return false
	}
	if set.ct >= CtBomb && set.ct > other.ct {
		return true
	}
	return false
}

func (set CardsSet) Dump() []int32 {
	t := make([]int32, len(set.seqs))
	copy(t, set.seqs)
	return t
}

func (set *CardsSet) Add(n []int32) {
	set.seqs = append(set.seqs, n...)
	set.refresh()
}

func (set *CardsSet) Del(n []int32) {
	set.seqs = base.SliceDel(set.seqs, n...)
	set.refresh()
}

func (set *CardsSet) continueCard(begin, size, count int) bool {
	for i := begin; i > begin-size; i-- {
		if i >= len(set.count) || i < 0 {
			return false
		}
		if set.count[i] != count {
			return false
		}
	}
	return true
}

func (set *CardsSet) SameColors() bool {
	colors := set.colors
	return colors[0]+colors[1]+colors[2]+colors[3]+colors[4] == 1
}

func (set *CardsSet) haveJokers() bool {
	return set.ContainSeq(0x4e, 0x4f)
}

func (set *CardsSet) ContainSeq(seqs ...int32) bool {
	return base.SliceContain(set.seqs, seqs...)
}

func (set *CardsSet) ContainNCard(n int) bool {
	return set.cnt[n] > 0
}

func (set *CardsSet) ContainChain(size, count int) bool {
	for lv, c := range set.count {
		if c == 0 {
			continue
		}
		if have := set.haveContinueCard(lv, size, count); have {
			return true
		}
	}
	return false
}

func (set *CardsSet) GetSeqCard() []int32 {
	return set.seqs
}
