package calc

/*
	Ai自动出牌算法，暂时未用到
*/

func (set *CardsSet) TimeOutOut(isFreeOut bool, last *CardsSet) (outCard *CardsSet) {

	if isFreeOut {
		len := len(set.seqs)
		if len > 0 {
			//默认出最小的一张
			outSeq := []int32{set.cards[len-1].Seq}
			outCard = NewCardsSet(outSeq)
		}
		return
	}

	find := set.Tips(last)
	if len(find) == 0 {
		return nil
	}
	return find[0]
}
