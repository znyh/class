package calc

func CheckHu(cardsIndex []int, guiIndex int) bool {
	return NewCardsSet(cardsIndex, nil, guiIndex).hu
}

func CheckHuCard(cardsIndex []int, cardIndex int, guiIndex int) bool {
	if cnt := CardsIndexCnt(cardsIndex); cnt%3 != 1 {
		return false
	}
	if cardIndex < 0 || cardIndex >= MAXINDEX {
		return false
	}
	count := make([]int, MAXINDEX)
	copy(count, cardsIndex)
	count[cardIndex]++

	return NewCardsSet(count, nil, guiIndex).hu
}

func CheckTing(cardsIndex []int, guiIndex int) bool {
	return len(NewCardsSet(cardsIndex, nil, guiIndex).ting) == 0
}

func TingTbl(cardsIndex []int, guiIndex int) (ting []int) {
	return NewCardsSet(cardsIndex, nil, guiIndex).ting
}

func OutTingTbl(cardsIndex []int, guiIndex int) (outTing map[int][]int) {
	outTing = map[int][]int{}
	count := make([]int, MAXINDEX)
	copy(count, cardsIndex)

	if cnt := CardsIndexCnt(cardsIndex); cnt%3 != 2 {
		return
	}

	for i := 0; i < MAXINDEX; i++ {
		ting := []int(nil)
		if count[i] > 0 {
			count[i]--
			ting = NewCardsSet(count, nil, guiIndex).ting
			count[i]++
		}
		if len(ting) > 0 {
			outTing[i] = ting
		}
	}
	return
}

func HuType(cardsIndex []int, weaves []TagWeaveItem, guiIndex int, target int) (ct int64) {
	set := NewCardsSet(cardsIndex, weaves, guiIndex)
	if !set.hu {
		return CtUnknown
	}

	if ok, super := set.CheckSuper7dui(); ok {
		switch super {
		case 3:
			ct |= CtSuper7DuiThree
		case 2:
			ct |= CtSuper7DuiDouble
		case 1:
			ct |= CtSuper7DuiSolo
		default:
			ct |= CtNormal7Dui
		}
	}

	if set.Check13Yao() {
		ct |= Ct13Yao
	}

	if set.CheckPengPengHu() {
		ct |= CtPengPeng
	} else if set.CheckKaCard(target) {
		ct |= CtKaZhang
	} else if set.CheckYiTiaoLong() {
		ct |= CtYitiaoLong
	}

	if set.CheckHuaLong() {
		ct |= CtHuaLong
	}

	if set.CheckQingYiSe() {
		ct |= CtQingYise
	} else if set.CheckZiYiSe() {
		ct |= CtHunYise
	} else if set.CheckZiYiSe() {
		ct |= CtZiYise
	}

	if ct == CtUnknown {
		ct = CtPingHU
	}
	return
}

func CardsIndexCnt(cbCardsIndex []int) int {
	cnt := 0
	for _, c := range cbCardsIndex {
		cnt += c
	}
	return cnt
}
