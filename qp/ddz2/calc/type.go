package calc

/*
	牌型算法
*/

const (
	CtUnknown        = 0  // 无效
	CtSolo           = 1  // 单张
	CtPair           = 2  // 对子
	CtThree          = 3  // 三张
	CtSoloChain      = 4  // 单顺
	CtPairChain      = 5  // 双顺
	CtThreeChain     = 6  // 三顺
	CtThreeSolo      = 7  // 三带一单
	CtThreePair      = 8  // 三带一对
	CtThreeSoloChain = 9  // 飞机+相同数量的单牌
	CtThreePairChain = 10 // 飞机+相同数量的对牌
	CtFour2          = 11 // 4带2张单或者1对
	CtFour4          = 12 // 4带2对
	CtBomb           = 13 // 炸弹
	CtJokers         = 14 // 大小王
)

var (
	checkCardType = make(map[int][]func(*CardsSet) bool)
)

func init() {
	checkCardType[1] = []func(*CardsSet) bool{checkSolo}
	checkCardType[2] = []func(*CardsSet) bool{checkPair, checkJokers}
	checkCardType[3] = []func(*CardsSet) bool{checkThree}
	checkCardType[4] = []func(*CardsSet) bool{checkThreeSolo, checkBomb}
	checkCardType[5] = []func(*CardsSet) bool{checkSoloChain, checkThreePair}
	checkCardType[6] = []func(*CardsSet) bool{checkSoloChain, checkPairChain, checkThreeChain, checkFour2}
	checkCardType[7] = []func(*CardsSet) bool{checkSoloChain}
	checkCardType[8] = []func(*CardsSet) bool{checkSoloChain, checkPairChain, checkThreeSoloChain, checkFour4}
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

//
func checkSolo(set *CardsSet) bool {
	return set.checkSolo()
}

func checkSoloChain(set *CardsSet) bool {
	return set.checkSoloChain()
}

func checkPair(set *CardsSet) bool {
	return set.checkPair()
}

func checkPairChain(set *CardsSet) bool {
	return set.checkPairChain()
}

func checkThree(set *CardsSet) bool {
	return set.checkThree()
}

func checkThreeChain(set *CardsSet) bool {
	return set.checkThreeChain()
}

func checkThreeSolo(set *CardsSet) bool {
	return set.checkThreeSolo()
}

func checkThreePair(set *CardsSet) bool {
	return set.checkThreePair()
}

func checkThreeSoloChain(set *CardsSet) bool {
	return set.checkThreeSoloChain()
}

func checkThreePairChain(set *CardsSet) bool {
	return set.checkThreePairChain()
}

func checkFour2(set *CardsSet) bool {
	return set.checkFour2()
}

func checkFour4(set *CardsSet) bool {
	return set.checkFour4()
}

func checkBomb(set *CardsSet) bool {
	return set.checkBomb()
}

func checkJokers(set *CardsSet) bool {
	return set.checkJokers()
}
