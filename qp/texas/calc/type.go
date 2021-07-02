package calc

const (
	CtHighCard        = 1
	CtOnePair         = 2
	CtTwoPair         = 3
	CtThree           = 4
	CtSoloChain       = 5
	CtFlush           = 6
	CtThreePair       = 7
	CtKind4           = 8
	CtFlushChain      = 9
	CtRoyalFlushChain = 10
)

var (
	checkCardType = []func(*CardSet) bool{
		checkRoyalFlushChain,
		checkFlushChain,
		checkKind4,
		checkThreePair,
		checkFlush,
		checkSoloChain,
		checkThree,
		checkTwoPair,
		checkOnePair,
		checkHighCard,
	}
)

func checkRoyalFlushChain(set *CardSet) bool {
	return set.CheckRoyalFlushChain()
}

func checkFlushChain(set *CardSet) bool {
	return set.CheckFlushChain()
}

func checkKind4(set *CardSet) bool {
	return set.CheckKind4()
}

func checkThreePair(set *CardSet) bool {
	return set.CheckThreePair()
}

func checkFlush(set *CardSet) bool {
	return set.CheckFlush()
}

func checkSoloChain(set *CardSet) bool {
	return set.CheckChain()
}

func checkThree(set *CardSet) bool {
	return set.CheckThree()
}

func checkTwoPair(set *CardSet) bool {
	return set.CheckTwoPair()
}

func checkOnePair(set *CardSet) bool {
	return set.CheckOnePair()
}

func checkHighCard(set *CardSet) bool {
	return set.CheckHighCard()
}
