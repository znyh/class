package calc

import (
	"github.com/znyh/class/qp/base"
)

/*
	Tips提示算法
*/

// Tips 提示
func (set *CardsSet) Tips(last *CardsSet) (find []*CardsSet) {
	last.calc()
	if last.ct == CtUnknown {
		return
	}

	switch last.ct {
	case CtJokers:
		return
	case CtBomb:
		return set.tipBomb(last.level)
	case CtSolo:
		return set.tipSolo(last.level)
	case CtPair:
		return set.tipPair(last.level)
	case CtThree:
		return set.tipThree(last.level)
	case CtThreeSolo:
		return set.tipThreeSolo(last.level)
	case CtThreePair:
		return set.tipThreePair(last.level)
	case CtSoloChain:
		return set.tipSoloChain(last.level, last.length)
	case CtPairChain:
		return set.tipPairChain(last.level, last.length)
	case CtThreeChain:
		return set.tipThreeChain(last.level, last.length)
	case CtThreeSoloChain:
		return set.tipThreeSoloChain(last.level, last.length)
	case CtThreePairChain:
		return set.tipThreePairChain(last.level, last.length)
	case CtFour2:
		return set.tipFour2(last.level)
	case CtFour4:
		return set.tipFour4(last.level)
	}
	return
}

func (set *CardsSet) HaveBigger(other *CardsSet) bool {
	return len(set.Tips(other)) > 0
}

func (set *CardsSet) getSeqsByLevel(lv int, cnt int) (seqs []int32, ok bool) {
	for _, v := range set.cards {
		if v.level == lv {
			seqs = append(seqs, v.Seq)
			if len(seqs) == cnt {
				ok = true
				return
			}
		}
	}
	return
}

func (set *CardsSet) haveContinueCard(begin, size, count int) bool {
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

func (set *CardsSet) biggerTip(ct int) (find []*CardsSet) {
	if ct < CtBomb {
		find = set.tipBomb(-1)
	}
	if set.ContainSeq(0x4e, 0x4f) {
		find = append(find, NewCardsSet([]int32{0x4e, 0x4f}))
	}
	return
}

func (set *CardsSet) tipBomb(bigLevel int) (find []*CardsSet) {
	for idx, v := range set.count {
		if v == 4 && idx > bigLevel {
			seqs, _ := set.getSeqsByLevel(idx, 4)
			find = append(find, NewCardsSet(seqs))
		}
	}
	return
}

func (set *CardsSet) tipSolo(bigLevel int) (find []*CardsSet) {
	for idx, v := range set.count {
		if v == 1 && idx > bigLevel {
			seqs, _ := set.getSeqsByLevel(idx, 1)
			find = append(find, NewCardsSet(seqs))
		}
	}
	for idx, v := range set.count {
		if v == 2 && idx > bigLevel {
			seqs, _ := set.getSeqsByLevel(idx, 1)
			find = append(find, NewCardsSet(seqs))
		}
	}
	for idx, v := range set.count {
		if v == 3 && idx > bigLevel {
			seqs, _ := set.getSeqsByLevel(idx, 1)
			find = append(find, NewCardsSet(seqs))
		}
	}
	if bf := set.biggerTip(CtSolo); bf != nil {
		find = append(find, bf...)
	}
	return
}

func (set *CardsSet) tipPair(bigLevel int) (find []*CardsSet) {
	for idx, v := range set.count {
		if v == 2 && idx > bigLevel {
			seqs, _ := set.getSeqsByLevel(idx, 2)
			find = append(find, NewCardsSet(seqs))
		}
	}
	for idx, v := range set.count {
		if v == 3 && idx > bigLevel {
			seqs, _ := set.getSeqsByLevel(idx, 2)
			find = append(find, NewCardsSet(seqs))
		}
	}
	if bf := set.biggerTip(CtPair); bf != nil {
		find = append(find, bf...)
	}
	return
}

func (set *CardsSet) tipThree(bigLevel int) (find []*CardsSet) {
	for idx, v := range set.count {
		if v == 3 && idx > bigLevel {
			seqs, _ := set.getSeqsByLevel(idx, 3)
			find = append(find, NewCardsSet(seqs))
		}
	}
	if bf := set.biggerTip(CtThree); bf != nil {
		find = append(find, bf...)
	}
	return
}

func (set *CardsSet) tipThreeSolo(bigLevel int) (find []*CardsSet) {
	for idx, v := range set.count {
		if v == 3 && idx > bigLevel {
			seqs, _ := set.getSeqsByLevel(idx, 3)
			xx := NewCardsSet(seqs)
			seqs2 := set.littleSolo([]int{idx})
			if seqs2 != nil {
				xx.Add(seqs2)
				find = append(find, xx)
			}
		}
	}
	if bf := set.biggerTip(CtThreeSolo); bf != nil {
		find = append(find, bf...)
	}
	return
}

func (set *CardsSet) tipThreePair(bigLevel int) (find []*CardsSet) {
	for idx, v := range set.count {
		if v == 3 && idx > bigLevel {
			seqs, _ := set.getSeqsByLevel(idx, 3)
			xx := NewCardsSet(seqs)
			seqs2 := set.littlePair([]int{idx})
			if seqs2 != nil {
				xx.Add(seqs2)
				find = append(find, xx)
			}
		}
	}
	if bf := set.biggerTip(CtThreePair); bf != nil {
		find = append(find, bf...)
	}
	return
}

func (set *CardsSet) tipSoloChain(bigLevel int, len int) (find []*CardsSet) {
	for start := bigLevel + 1; start < 13; start++ {
		if set.haveContinueCard(start, len, 1) {
			s := NewCardsSet([]int32{})
			for i := start; i > (start - len); i-- {
				seqs, _ := set.getSeqsByLevel(i, 1)
				s.Add(seqs)
			}
			find = append(find, s)
		}
	}
	if bf := set.biggerTip(CtSoloChain); bf != nil {
		find = append(find, bf...)
	}
	return
}

func (set *CardsSet) tipPairChain(bigLevel int, len int) (find []*CardsSet) {
	for start := bigLevel + 1; start < 13; start++ {
		if set.haveContinueCard(start, len, 2) {
			s := NewCardsSet([]int32{})
			for i := start; i > (start - len); i-- {
				seqs, _ := set.getSeqsByLevel(i, 2)
				s.Add(seqs)
			}
			find = append(find, s)
		}
	}
	if bf := set.biggerTip(CtPairChain); bf != nil {
		find = append(find, bf...)
	}
	return
}

func (set *CardsSet) tipThreeChain(bigLevel int, len int) (find []*CardsSet) {
	for start := bigLevel + 1; start < 13; start++ {
		if set.haveContinueCard(start, len, 3) {
			s := NewCardsSet([]int32{})
			for i := start; i > (start - len); i-- {
				seqs, _ := set.getSeqsByLevel(i, 3)
				s.Add(seqs)
			}
			find = append(find, s)
		}
	}
	if bf := set.biggerTip(CtThreeChain); bf != nil {
		find = append(find, bf...)
	}
	return
}

func (set *CardsSet) tipThreeSoloChain(bigLevel int, size int) (find []*CardsSet) {
	for start := bigLevel + 1; start < 13; start++ {
		if set.haveContinueCard(start, size, 3) {
			set1 := NewCardsSet([]int32{})
			for i := start; i > (start - size); i-- {
				seqs, _ := set.getSeqsByLevel(i, 3)
				set1.Add(seqs)
			}

			left := base.SliceDel(base.SliceCopy(set.seqs), set1.seqs...)

			if containBombLimit {
				for _, seq := range left {
					lv := cardLevel(seq)
					if set.count[lv] == 4 {
						left = base.SliceDel(left, seq)
					}
				}
			}

			if dfs := base.Dfs(left, size); len(dfs) != 0 {
				for _, v := range dfs {
					if !base.SliceContain(v, 0x4e, 0x4f) {
						newSeqs := []int32(nil)
						newSeqs = append(newSeqs, set1.seqs...)
						newSeqs = append(newSeqs, v...)
						find = append(find, NewCardsSet(newSeqs))
					}
				}
			}

			//notLevel := []int(nil)
			//for zz := 0; zz < len; zz++ {
			//	notLevel = append(notLevel, start-zz)
			//}
			//for soloIdx := 0; soloIdx < len; soloIdx++ {
			//	soloX := set.littleSolo(notLevel)
			//	if soloX == nil {
			//		return
			//	}
			//	set1.Add(soloX)
			//	notLevel = append(notLevel, cardLevel(soloX[0]))
			//}
			//find = append(find, set1)
		}
	}
	if bf := set.biggerTip(CtThreeSoloChain); bf != nil {
		find = append(find, bf...)
	}
	return
}

func (set *CardsSet) tipThreePairChain(bigLevel int, len int) (find []*CardsSet) {
	for start := bigLevel + 1; start < 13; start++ {
		if set.haveContinueCard(start, len, 3) {
			s := NewCardsSet([]int32{})
			for i := start; i > (start - len); i-- {
				seqs, _ := set.getSeqsByLevel(i, 3)
				s.Add(seqs)
			}

			notLevel := []int(nil)
			for zz := 0; zz < len; zz++ {
				notLevel = append(notLevel, start-zz)
			}
			for soloIdx := 0; soloIdx < len; soloIdx++ {
				soloX := set.littlePair(notLevel)
				if soloX == nil {
					return
				}
				s.Add(soloX)
				notLevel = append(notLevel, cardLevel(soloX[0]))
			}
			find = append(find, s)
		}
	}
	if bf := set.biggerTip(CtThreePairChain); bf != nil {
		find = append(find, bf...)
	}
	return
}

func (set *CardsSet) tipFour2(bigLevel int) (find []*CardsSet) {
	for idx, v := range set.count {
		if v == 4 && idx > bigLevel {
			seqs, _ := set.getSeqsByLevel(idx, 4)
			xx := NewCardsSet(seqs)

			seqs2 := set.littlePair([]int{idx})
			if seqs2 != nil {
				xx.Add(seqs2)
				find = append(find, xx)
			} else {
				seq3 := set.littleSolo([]int{idx})
				if seq3 == nil || len(seq3) != 1 {
					return
				}
				idx3 := cardLevel(seq3[0])
				seq4 := set.littleSolo([]int{idx, idx3})
				if seq4 == nil {
					return
				}
				xx.Add(seq3)
				xx.Add(seq4)
				find = append(find, xx)
			}
		}
	}
	if bf := set.biggerTip(CtFour2); bf != nil {
		find = append(find, bf...)
	}
	return
}

func (set *CardsSet) tipFour4(bigLevel int) (find []*CardsSet) {
	for idx, v := range set.count {
		if v == 4 && idx > bigLevel {
			seqs, _ := set.getSeqsByLevel(idx, 4)
			xx := NewCardsSet(seqs)

			seqs2 := set.littlePair([]int{idx})
			if seqs2 == nil {
				return
			}
			idx2 := cardLevel(seqs2[0])
			seqs3 := set.littlePair([]int{idx, idx2})
			if seqs3 == nil {
				return
			}

			xx.Add(seqs2)
			xx.Add(seqs3)
			find = append(find, xx)
		}
	}
	if bf := set.biggerTip(CtFour4); bf != nil {
		find = append(find, bf...)
	}
	return
}

// 找一个最小且不为notLevel单张
func (set *CardsSet) littleSolo(notLevel []int) []int32 {
	for idx, v := range set.count {
		if v == 0 {
			continue
		}

		find := false
		for _, v2 := range notLevel {
			if v2 == idx {
				find = true
				break
			}
		}
		if find {
			continue
		}

		seqs, _ := set.getSeqsByLevel(idx, 1)
		return seqs
	}
	return nil
}

// 找一个最小且不为notLevel对子
func (set *CardsSet) littlePair(notLevel []int) []int32 {
	for idx, v := range set.count {
		if v < 2 {
			continue
		}

		find := false
		for _, v2 := range notLevel {
			if v2 == idx {
				find = true
				break
			}
		}
		if find {
			continue
		}

		seqs, _ := set.getSeqsByLevel(idx, 2)
		return seqs
	}
	return nil
}
