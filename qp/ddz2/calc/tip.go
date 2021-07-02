package calc

func (set *CardsSet) haveJokers() bool {
	return set.haveSequence(0x4e, 0x4f)
}

func (set *CardsSet) haveSequence(sequence ...int32) bool {
	for _, seq := range sequence {
		find := false
		for _, seq2 := range set.seqs {
			if seq2 == seq {
				find = true
			}
		}
		if !find {
			return false
		}
	}
	return true
}

func (set *CardsSet) haveContinue(begin, size, count int) (bool, int) {
	need := 0
	for i := begin; i > begin-size; i-- {
		if i >= Card2Idx() || i < 0 {
			return false, 0
		}
		if set.count[i] <= count {
			need += count - set.count[i]
		}
	}
	return set.guiNum >= need, need
}

func (set *CardsSet) continueCard(begin, size, count int) (bool, int) {
	need := 0
	for i := begin; i > begin-size; i-- {
		if i >= Card2Idx() || i < 0 {
			return false, 0
		}
		if set.count[i] > count {
			return false, 0
		}
		if set.count[i] <= count {
			need += count - set.count[i]
		}
	}
	return set.guiNum >= need, need
}
