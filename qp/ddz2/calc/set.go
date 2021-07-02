package calc

import (
	"sort"
)

type CardsSet struct {
	seqs     []int32              // 不会排序
	guiSeqs  []int32              // 赖子序列
	size     int                  //
	guiNum   int                  //
	cards    []int32              //按牌值花色降序排序
	guiIndex []int                //赖子的cardsIndex
	count    [MAXCARD]int         //去掉赖子后的cardsIndex
	countSeq [MAXCARD][]int32     //cardsIndex => cardsData
	cnt      map[int]int          //key:{1,2,3,4} => num:{1,2,3,4}
	cts      map[int32][]cardType // key:cardType
}

type cardType struct {
	ct        int
	level     int
	continues int
}

func newType(ct int, level int, continues int) cardType {
	return cardType{
		ct:        ct,
		level:     level,
		continues: continues,
	}
}

func NewCardsSet(seqs []int32, guiseqs []int32) *CardsSet {
	set := &CardsSet{}

	if !ValidCardsData(seqs) {
		return set
	}

	if !ValidCardsData(guiseqs) {
		return set
	}

	set.seqs = make([]int32, len(seqs))
	copy(set.seqs, seqs)

	set.guiSeqs = make([]int32, len(guiseqs))
	copy(set.guiSeqs, guiseqs)

	set.refresh()

	return set
}

func (set *CardsSet) refresh() {
	set.size = len(set.seqs)
	set.count = [MAXCARD]int{}
	set.countSeq = [MAXCARD][]int32{}
	set.cnt = map[int]int{}
	set.cts = make(map[int32][]cardType)

	for _, seq := range set.seqs {
		if SliceContain(set.guiSeqs, seq) {
			set.guiNum++
			set.guiIndex = append(set.guiIndex, ToCardIndex(seq))
			//set.count[MAXCARD-1]++
			set.countSeq[MAXCARD-1] = append(set.countSeq[MAXCARD-1], seq)
		} else {
			idx := ToCardIndex(seq)
			set.count[idx]++
			set.countSeq[idx] = append(set.countSeq[idx], seq)
		}
	}

	for k := ToCardIndex(0x4e); k >= 0; k-- {
		sequences := set.countSeq[k]
		sort.Slice(sequences, func(i, j int) bool {
			return sequences[i] > sequences[j]
		})
		set.cards = append(set.cards, sequences...)
	}

	for idx, c := range set.count {
		if c > 0 && idx < MAXCARD {
			set.cnt[c]++
		}
	}

	set.calc()
}

func (set *CardsSet) calc() {
	for _, v := range checkCardType[set.size] {
		v(set)
	}
}

func (set *CardsSet) checkSolo() bool {
	if set.size == 1 {
		ct := newType(CtSolo, ToCardIndex(set.seqs[0]), 0)
		set.cts[CtSolo] = append(set.cts[CtSolo], ct)
		return true
	}
	return false
}

func (set *CardsSet) checkPair() bool {
	if set.size != 2 {
		return false
	}
	for idx := Card2Idx(); idx >= 0; idx-- {
		c := set.count[idx]
		if c > 0 && c+set.guiNum == 2 {
			ct := newType(CtPair, idx, 0)
			set.cts[CtPair] = append(set.cts[CtPair], ct)
		}
	}
	if set.guiNum == 2 {
		ct := newType(CtPair, set.guiIndex[0], 0)
		set.cts[CtPair] = append(set.cts[CtPair], ct)
	}
	return len(set.cts[CtPair]) > 0
}

func (set *CardsSet) checkThree() bool {
	if set.size != 3 {
		return false
	}
	for idx := Card2Idx(); idx >= 0; idx-- {
		c := set.count[idx]
		if c > 0 && c+set.guiNum == 3 {
			ct := newType(CtThree, idx, 0)
			set.cts[CtThree] = append(set.cts[CtThree], ct)
		}
	}
	if set.guiNum == 3 {
		ct := newType(CtThree, set.guiIndex[0], 0)
		set.cts[CtThree] = append(set.cts[CtThree], ct)
	}
	return len(set.cts[CtThree]) > 0
}

func (set *CardsSet) checkSoloChain() bool {
	if set.size < 5 || set.size >= 13 {
		return false
	}
	count := 1
	shouldLen := set.size / count
	for end := CardAIdx(); end >= 0; end-- {
		if set.count[end] > count {
			return false
		}
		ok, need := set.continueCard(end, shouldLen, count)
		if ok && set.guiNum-need == 0 {
			ct := newType(CtSoloChain, end, shouldLen)
			set.cts[CtSoloChain] = append(set.cts[CtSoloChain], ct)
		}
	}
	return len(set.cts[CtSoloChain]) > 0
}

func (set *CardsSet) checkPairChain() bool {
	if set.size%2 != 0 || set.size < 6 && set.size > 20 {
		return false
	}
	count := 2
	shouldLen := set.size / count
	for end := CardAIdx(); end >= 0; end-- {
		if set.count[end] > count {
			return false
		}
		ok, need := set.continueCard(end, shouldLen, count)
		if ok && set.guiNum-need == 0 {
			ct := newType(CtPairChain, end, shouldLen)
			set.cts[CtPairChain] = append(set.cts[CtPairChain], ct)
		}
	}
	return len(set.cts[CtPairChain]) > 0
}

func (set *CardsSet) checkThreeChain() bool {
	if set.size%3 != 0 || set.size < 6 && set.size > 18 {
		return false
	}
	count := 3
	shouldLen := set.size / count
	for end := CardAIdx(); end >= 0; end-- {
		if set.count[end] > count {
			return false
		}
		ok, need := set.continueCard(end, shouldLen, count)
		if ok && set.guiNum-need == 0 {
			ct := newType(CtThreeChain, end, shouldLen)
			set.cts[CtThreeChain] = append(set.cts[CtThreeChain], ct)
		}
	}
	return len(set.cts[CtThreeChain]) > 0
}

func (set *CardsSet) checkThreeSolo() bool {
	if set.size != 4 {
		return false
	}
	if set.guiNum+1*set.cnt[1]+2*set.cnt[2]+3*set.cnt[3] != 4 {
		return false
	}
	for i := 0; i <= Card2Idx(); i++ {
		if set.count[i] <= 0 {
			continue
		}
		left := set.guiNum - (3 - set.count[i])
		if left >= 0 && left <= 1 {
			ct := newType(CtThreeSolo, i, 0)
			set.cts[CtThreeSolo] = append(set.cts[CtThreeSolo], ct)
		}
	}
	if set.guiNum == 3 {
		ct := newType(CtThreeSolo, set.guiIndex[0], 0)
		set.cts[CtThreeSolo] = append(set.cts[CtThreeSolo], ct)
	}
	return len(set.cts[CtThreeSolo]) > 0
}

func (set *CardsSet) checkThreePair() bool {
	if set.size != 5 {
		return false
	}
	if set.guiNum+1*set.cnt[1]+2*set.cnt[2]+3*set.cnt[3] != 5 {
		return false
	}
	for i := 0; i <= Card2Idx(); i++ {
		if set.count[i] <= 0 {
			continue
		}
		left := set.guiNum - (3 - set.count[i])
		if left < 0 || left > 2 {
			continue
		}
		for j := 0; i != j && j <= Card2Idx(); j++ {
			if set.count[j] > 0 && left+set.count[j] == 2 {
				ct := newType(CtThreePair, i, 0)
				set.cts[CtThreePair] = append(set.cts[CtThreePair], ct)
			}
		}
		if set.count[i]+set.guiNum == set.size {
			ct := newType(CtThreePair, i, 0)
			set.cts[CtThreePair] = append(set.cts[CtThreePair], ct)
		}
	}
	if (set.guiNum == 3 && set.cnt[2] == 1) || (set.guiNum == 4 && set.cnt[1] == 1) {
		ct := newType(CtThreePair, set.guiIndex[0], 0)
		set.cts[CtThreePair] = append(set.cts[CtThreePair], ct)
	}
	return len(set.cts[CtThreePair]) > 0
}

func (set *CardsSet) checkThreeSoloChain() bool {
	if set.size%4 != 0 || set.size < 8 || set.size > 20 {
		return false
	}

	count := 3
	shouldLen := set.size / 4
	for end := CardAIdx(); end >= 0; end-- {
		if set.count[end] < 0 {
			continue
		}
		ok, need := set.haveContinue(end, shouldLen, count)
		if ok && need <= set.guiNum {
			ct := newType(CtThreeSoloChain, end, shouldLen)
			set.cts[CtThreeSoloChain] = append(set.cts[CtThreeSoloChain], ct)
		}
	}
	return len(set.cts[CtThreeSoloChain]) > 0
}

func (set *CardsSet) checkThreePairChain() bool {
	if set.size%5 != 0 || set.size < 10 || set.size > 20 {
		return false
	}
	//shouldLen := length / 5
	return len(set.cts[CtThreePairChain]) > 0
}

func (set *CardsSet) checkFour2() bool {
	if set.size != 6 || set.haveJokers() {
		return false
	}
	for i := 0; i <= Card2Idx(); i++ {
		if set.count[i] <= 0 {
			continue
		}
		left := set.guiNum - (4 - set.count[i])
		if left >= 0 {
			ct := newType(CtFour2, i, 0)
			set.cts[CtFour2] = append(set.cts[CtFour2], ct)
		}
	}
	if set.guiNum == 4 {
		ct := newType(CtFour2, set.guiIndex[0], 0)
		set.cts[CtFour2] = append(set.cts[CtFour2], ct)
	}
	return len(set.cts[CtFour2]) > 0
}

func (set *CardsSet) checkFour4() bool {
	if set.size != 8 {
		return false
	}
	for i := 0; i <= Card2Idx(); i++ {
		if set.count[i] <= 0 {
			continue
		}
		cnt2 := 0
		left := set.guiNum - (4 - set.count[i])
		if left < 0 {
			continue
		}
		for j := 0; i != j && j <= Card2Idx(); j++ {
			if set.count[j] > 0 && left+set.count[j] == 2 {
				cnt2++
				left -= 2 - set.count[j]
			}
		}
		if left >= 0 && left+2*cnt2 == 4 {
			ct := newType(CtFour4, i, 0)
			set.cts[CtFour4] = append(set.cts[CtFour4], ct)
		}
	}
	if set.guiNum == 4 && set.cnt[2] == 2 {
		ct := newType(CtFour4, set.guiIndex[0], 0)
		set.cts[CtFour4] = append(set.cts[CtFour4], ct)
	}

	return len(set.cts[CtFour4]) > 0
}

func (set *CardsSet) checkBomb() bool {
	if set.size == 4 {
		for idx := 0; idx <= Card2Idx(); idx++ {
			c := set.count[idx]
			if c > 0 && c+set.guiNum == 4 {
				ct := newType(CtBomb, idx, 0)
				set.cts[CtBomb] = append(set.cts[CtBomb], ct)
			}
		}
		if set.guiNum == 4 {
			ct := newType(CtBomb, set.guiIndex[0], 0)
			set.cts[CtBomb] = append(set.cts[CtBomb], ct)
		}
	}
	return len(set.cts[CtBomb]) > 0
}

func (set *CardsSet) checkJokers() bool {
	if set.size == 2 && set.haveJokers() {
		ct := newType(CtJokers, ToCardIndex(0x4f), 0)
		set.cts[CtJokers] = append(set.cts[CtJokers], ct)
		return true
	}
	return false
}
