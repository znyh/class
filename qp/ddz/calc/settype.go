package calc

/*
	牌型算法
*/

const (
	CtUnknown        = iota // 无效
	CtSolo                  // 单张
	CtPair                  // 对子
	CtThree                 // 三张
	CtSoloChain             // 单顺
	CtPairChain             // 双顺
	CtThreeChain            // 三顺
	CtThreeSolo             // 三带一单
	CtThreePair             // 三带一对
	CtThreeSoloChain        // 飞机+相同数量的单牌
	CtThreePairChain        // 飞机+相同数量的对牌
	CtFour2                 // 4带2张单或者1对
	CtFour4                 // 4带2对
	CtBomb                  // 炸弹
	CtJokers                // 大小王
)

var (
	checkCardType = make(map[int][](func(*CardsSet) bool))
)

func init() {
	checkCardType[1] = []func(*CardsSet) bool{checkSolo}
	checkCardType[2] = []func(*CardsSet) bool{checkPair, checkJokers}
	checkCardType[3] = []func(*CardsSet) bool{checkThree}
	checkCardType[4] = []func(*CardsSet) bool{checkThreeSolo, checkBomb}
	checkCardType[5] = []func(*CardsSet) bool{checkSoloChain, checkThreePair}
	checkCardType[6] = []func(*CardsSet) bool{checkSoloChain, checkPairChain, checkThreeChain, checkFour2}
	checkCardType[7] = []func(*CardsSet) bool{checkSoloChain}
	checkCardType[8] = []func(*CardsSet) bool{checkFour4, checkSoloChain, checkPairChain, checkThreeSoloChain}
	checkCardType[9] = []func(*CardsSet) bool{checkSoloChain, checkThreeChain}
	checkCardType[10] = []func(*CardsSet) bool{checkSoloChain, checkPairChain, checkThreePairChain}
	checkCardType[11] = []func(*CardsSet) bool{checkSoloChain}
	checkCardType[12] = []func(*CardsSet) bool{checkSoloChain, checkPairChain, checkThreeChain, checkThreeSoloChain}
	checkCardType[14] = []func(*CardsSet) bool{checkPairChain}
	checkCardType[15] = []func(*CardsSet) bool{checkThreeChain, checkThreePairChain}
	checkCardType[16] = []func(*CardsSet) bool{checkPairChain, checkThreeSoloChain}
	checkCardType[18] = []func(*CardsSet) bool{checkPairChain, checkThreeChain}
	checkCardType[20] = []func(*CardsSet) bool{checkPairChain, checkThreeSoloChain, checkThreePairChain}
}

func checkSolo(set *CardsSet) bool {
	if len(set.cards) == 1 {
		set.ct = CtSolo
		set.level = set.cards[0].level
		return true
	}
	return false
}

func checkSoloChain(set *CardsSet) bool {
	length := len(set.cards)
	if length >= 5 && length < 13 {
		level := set.cards[0].level
		if level > number2Level(1) {
			return false
		}

		shouldLen := length
		if set.continueCard(level, shouldLen, 1) {
			set.ct = CtSoloChain
			set.length = shouldLen
			set.level = level
			return true
		}
	}
	return false
}

func checkPair(set *CardsSet) bool {
	if len(set.cards) == 2 && set.cards[0].level == set.cards[1].level {
		set.ct = CtPair
		set.level = set.cards[0].level
		return true
	}
	return false
}

func checkPairChain(set *CardsSet) bool {
	length := len(set.cards)
	if length%2 == 0 && length >= 6 && length <= 20 {
		level := set.cards[0].level
		if level > number2Level(1) {
			return false
		}

		shouldLen := length / 2
		if set.continueCard(level, shouldLen, 2) {
			set.ct = CtPairChain
			set.length = shouldLen
			set.level = level
			return true
		}
	}
	return false
}

func checkThree(set *CardsSet) bool {
	if len(set.cards) == 3 &&
		set.cards[0].level == set.cards[1].level &&
		set.cards[0].level == set.cards[2].level {
		set.ct = CtThree
		set.level = set.cards[0].level
		return true
	}
	return false
}

func checkThreeChain(set *CardsSet) bool {
	length := len(set.cards)
	if length%3 == 0 && length >= 6 && length <= 18 {
		level := set.cards[0].level
		if level > number2Level(1) {
			return false
		}

		shouldLen := length / 3
		if set.continueCard(level, shouldLen, 3) {
			set.ct = CtThreeChain
			set.length = shouldLen
			set.level = level
			return true
		}
	}
	return false
}

func checkThreeSolo(set *CardsSet) bool {
	if len(set.cards) == 4 {
		cnt1 := 0
		cnt3 := 0
		level := 0
		for idx, v := range set.count {
			if v == 1 {
				cnt1++
			} else if v == 3 {
				cnt3++
				level = idx
			}
		}
		if cnt1 == 1 && cnt3 == 1 {
			set.ct = CtThreeSolo
			set.level = level
			return true
		}
	}
	return false
}

func checkThreePair(set *CardsSet) bool {
	if len(set.cards) == 5 {
		cnt2 := 0
		cnt3 := 0
		level := 0
		for idx, v := range set.count {
			if v == 2 {
				cnt2++
			} else if v == 3 {
				cnt3++
				level = idx
			}
		}
		if cnt2 == 1 && cnt3 == 1 {
			set.ct = CtThreePair
			set.level = level
			return true
		}
	}
	return false
}

func checkThreeSoloChain(set *CardsSet) bool {
	if set.haveJokers() {
		return false
	}

	if containBombLimit {
		if set.cnt[4] > 0 {
			return false
		}
	}

	length := len(set.cards)
	if length%4 == 0 && length >= 8 && length <= 20 {
		shouldLen := length / 4

		for idx, v := range set.count {
			if v >= 3 {
				if set.haveContinueCard(idx+shouldLen-1, shouldLen, 3) {
					set.ct = CtThreeSoloChain
					set.length = shouldLen
					set.level = idx + shouldLen - 1
					return true
				}
				break
			}
		}
	}
	return false
}

func checkThreePairChain(set *CardsSet) bool {
	if set.haveJokers() || set.cnt[4] > 0 {
		return false
	}

	length := len(set.cards)
	if length%5 == 0 && length >= 10 && length <= 20 {
		shouldLen := length / 5
		if set.cnt[3] != shouldLen || set.cnt[2] != shouldLen {
			return false
		}

		for idx, v := range set.count {
			if v == 3 {
				if set.continueCard(idx+shouldLen-1, shouldLen, 3) {
					set.ct = CtThreePairChain
					set.length = shouldLen
					set.level = idx + shouldLen - 1
					return true
				}
				break
			}
		}
	}
	return false
}

func checkFour2(set *CardsSet) bool {
	if set.haveJokers() {
		return false
	}

	if len(set.cards) == 6 {
		cnt4 := 0
		level := 0
		for idx, v := range set.count {
			if v == 4 {
				cnt4++
				level = idx
				break
			}
		}
		if cnt4 == 1 {
			set.ct = CtFour2
			set.level = level
			return true
		}
	}
	return false
}

func checkFour4(set *CardsSet) bool {

	if containBombLimit {
		if set.cnt[4] >= 2 {
			return false
		}
	} else {
		if specialFour4(set) {
			return true
		}
	}

	if len(set.cards) == 8 {
		cnt2 := 0
		cnt4 := 0
		level := 0
		for idx, v := range set.count {
			if v == 4 {
				cnt4++
				level = idx
			} else if v == 2 {
				cnt2++
			}
		}
		if cnt4 == 1 && cnt2 == 2 {
			set.ct = CtFour4
			set.level = level
			return true
		}
	}
	return false
}

func specialFour4(set *CardsSet) bool {
	if len(set.cards) == 8 && set.cnt[4] == 2 {
		c := []int(nil)
		for idx, v := range set.count {
			if v == 4 {
				c = append(c, idx)
			}
		}
		if c[0] > c[1] {
			c[0], c[1] = c[1], c[0]
		}
		set.ct = CtFour4
		set.level = c[1]
		return true
	}
	return false
}

func checkBomb(set *CardsSet) bool {
	if len(set.cards) == 4 {
		for idx, v := range set.count {
			if v == 4 {
				set.ct = CtBomb
				set.level = idx
				return true
			}
		}
	}
	return false
}

func checkJokers(set *CardsSet) bool {
	if len(set.cards) == 2 && set.cards[0].Seq == 0x4f && set.cards[1].Seq == 0x4e {
		set.ct = CtJokers
		return true
	}
	return false
}
