package calc

import (
	"sort"
)

type CardSet struct {
	seqs     []int32 // 不会排序
	cards    []Card
	count    [_MaxCard]int
	color    map[int]int
	colorMap map[int][]int32
	outs     []int32
	level    int
	ct       int
}

func NewCardSet(seqs []int32) *CardSet {
	set := &CardSet{}

	set.seqs = make([]int32, len(seqs))
	copy(set.seqs, seqs)

	set.refresh()

	return set
}

func (set *CardSet) refresh() {
	set.cards = make([]Card, len(set.seqs))
	set.count = [_MaxCard]int{}
	set.color = map[int]int{}
	set.colorMap = map[int][]int32{}

	for i := 0; i < len(set.seqs); i++ {
		set.cards[i].Seq = set.seqs[i]
		set.cards[i].calc()
		set.count[set.cards[i].level]++
		set.color[set.cards[i].color] = 1
		set.colorMap[set.cards[i].color] = append(set.colorMap[set.cards[i].color], set.seqs[i])
	}

	sort.Slice(set.cards, func(i, j int) bool {
		return set.cards[i].level > set.cards[j].level
	})

	sort.Slice(set.cards, func(i, j int) bool {
		if set.cards[i].level == set.cards[j].level {
			return set.cards[i].color > set.cards[j].color
		}
		return false
	})

	for k, _ := range set.colorMap {
		sort.Slice(set.colorMap[k], func(i, j int) bool {
			return set.colorMap[k][i] > set.colorMap[k][j]
		})
	}

	set.calc()
}

func (set *CardSet) calc() {
	for _, v := range checkCardType {
		if v(set) {
			return
		}
	}
}

func (set *CardSet) CheckRoyalFlushChain() bool {
	const (
		_Length = 5
		_Color  = 3
		_Start  = _MaxCard - 1
	)
	if set.haveFlushContinueCard(_Start, _Length, 1, _Color) {
		out := []int32(nil)
		for i := _Start; i > _Start-_Length; i-- {
			_, seq := set.getSeqByFlush(i, _Color)
			out = append(out, seq)
		}
		set.ct = CtRoyalFlushChain
		set.outs = out
		set.level = _Start
		return true
	}
	return false
}

func (set *CardSet) CheckFlushChain() bool {
	const _Length = 5
	for color, arr := range set.colorMap {
		if len(arr) >= 5 {
			for start := _MaxCard - 1; start >= 3; start-- {
				if set.haveFlushContinueCard(start, _Length, 1, color) {
					out := []int32(nil)
					for i := start; i > start-_Length; i-- {
						_, seq := set.getSeqByFlush(i, color)
						out = append(out, seq)
					}
					set.ct = CtFlushChain
					set.outs = out
					set.level = start
					return true
				}
			}
		}
	}
	return false
}

func (set *CardSet) CheckKind4() bool {
	for level := _MaxCard - 1; level >= 2; level-- {
		if set.count[level] == 4 {
			var out = []int32(nil)
			_, seq4 := set.getSeqByLevel(level, 4)
			out = append(out, seq4...)

			for start := _MaxCard - 1; start >= 3; start-- {
				if start != level && set.count[start] >= 1 {
					_, seq1 := set.getSeqByLevel(start, 1)
					out = append(out, seq1...)

					set.ct = CtKind4
					set.level = level
					set.outs = out
					return true
				}
			}
		}
	}
	return false
}

func (set *CardSet) CheckThreePair() bool {
	for level := _MaxCard - 1; level >= 2; level-- {
		if set.count[level] == 3 {
			var out = []int32(nil)
			_, seq3 := set.getSeqByLevel(level, 3)
			out = append(out, seq3...)

			for start := _MaxCard - 1; start >= 3; start-- {
				if start != level && set.count[start] >= 2 {
					_, seq2 := set.getSeqByLevel(start, 2)
					out = append(out, seq2...)

					set.ct = CtThreePair
					set.level = level
					set.outs = out
					return true
				}
			}
		}
	}
	return false
}

func (set *CardSet) CheckFlush() bool {
	for _, out := range set.colorMap {
		if len(out) < 5 {
			continue
		}
		set.ct = CtFlush
		set.level = cardLevel(out[0])
		set.outs = out
		return true
	}
	return false
}

func (set *CardSet) CheckChain() bool {
	const _Length = 5
	for start := _MaxCard - 1; start >= 3; start-- {
		if set.haveContinueCard(start, _Length, 1) {
			out := []int32(nil)
			for i := start; i > start-_Length; i-- {
				_, seq := set.getSeqByLevel(i, 1)
				out = append(out, seq...)
			}
			set.ct = CtSoloChain
			set.outs = out
			set.level = start
			return true
		}
	}
	return false
}

func (set *CardSet) CheckThree() bool {
	for level := _MaxCard - 1; level >= 2; level-- {
		if set.count[level] == 3 {
			var out = []int32(nil)
			var cnt1 = 0
			_, seq4 := set.getSeqByLevel(level, 3)
			out = append(out, seq4...)

			for start := _MaxCard - 1; start >= 3 && cnt1 < 2; start-- {
				if start != level && set.count[start] <= 1 {
					_, seq1 := set.getSeqByLevel(start, 1)
					out = append(out, seq1...)
					cnt1++
				}
			}
			if cnt1 == 2 {
				set.ct = CtThree
				set.level = level
				set.outs = out
				return true
			}
		}
	}
	return false
}

func (set *CardSet) CheckTwoPair() bool {
	var (
		cnt2     = 0
		maxLevel = []int32(nil)
		out      = []int32(nil)
	)
	for level := _MaxCard - 1; level >= 2; level-- {
		if set.count[level] == 2 {
			cnt2++
			maxLevel = append(maxLevel, int32(level))
			_, seq2 := set.getSeqByLevel(level, 2)
			out = append(out, seq2...)

			if cnt2 == 2 {
				for start := _MaxCard - 1; start >= 3; start-- {
					if !sliceContain(maxLevel, int32(start)) && set.count[start] >= 1 {
						_, seq1 := set.getSeqByLevel(start, 1)
						out = append(out, seq1...)

						set.ct = CtTwoPair
						set.level = int(maxLevel[0])
						set.outs = out
						return true
					}
				}
			}
		}

	}
	return false
}

func (set *CardSet) CheckOnePair() bool {
	for level := _MaxCard - 1; level >= 2; level-- {
		if set.count[level] == 2 {
			var out = []int32(nil)
			var cnt1 = 0

			_, seq2 := set.getSeqByLevel(level, 2)
			out = append(out, seq2...)

			for start := _MaxCard - 1; start >= 3 && cnt1 < 3; start-- {
				if start != level && set.count[start] <= 1 {
					_, seq1 := set.getSeqByLevel(start, 1)
					out = append(out, seq1...)
					cnt1++
				}
			}
			if cnt1 == 3 {
				set.ct = CtOnePair
				set.level = level
				set.outs = out
				return true
			}
		}
	}
	return false
}

func (set *CardSet) CheckHighCard() bool {
	var (
		cnt1     = 0
		maxLevel = 0
		out      = []int32(nil)
	)
	for level := _MaxCard - 1; level >= 3 && cnt1 < 5; level-- {
		if set.count[level] <= 1 {
			_, seq1 := set.getSeqByLevel(level, 1)
			out = append(out, seq1...)
			cnt1++
		}
		if cnt1 == 1 {
			maxLevel = level
		} else if cnt1 == 5 {
			set.ct = CtHighCard
			set.level = maxLevel
			set.outs = out
			return true
		}
	}
	return false
}

func (set *CardSet) haveContinueCard(begin, size, count int) bool {
	for i := begin; i > begin-size; i-- {
		if i >= len(set.count) || i < 0 {
			return false
		}
		if set.count[i] < count {
			return false
		}
	}
	return true
}

func (set *CardSet) haveFlushContinueCard(begin, size, count, color int) bool {
	for i := begin; i > begin-size; i-- {
		if set.count[i] < count {
			return false
		}
		if ok, _ := set.getSeqByFlush(i, color); !ok {
			return false
		}
	}
	return true
}

func (set *CardSet) getSeqByFlush(level int, color int) (ok bool, seq int32) {
	for _, v := range set.cards {
		if v.level == level && v.color == color {
			return true, v.Seq
		}
	}
	return false, -1
}

func (set *CardSet) getSeqByLevel(level int, count int) (ok bool, seqs []int32) {
	for _, v := range set.cards {
		if v.level == level {
			seqs = append(seqs, v.Seq)
			if len(seqs) == count {
				ok = true
				return
			}
		}
	}
	return
}

func (set *CardSet) ContainSeq(seqs ...int32) bool {
	return SliceContain(set.seqs, seqs...)
}
