package calc

import (
	"github.com/znyh/class/qp/base"
)

const (
	_MaxLv        = 15 // '3' -> joker
	_UnknownLevel = 0  //无效牌
)

type ShuffleConfig struct {
	//首次发牌规则洗牌
	UseFirstRule  bool            //是否使用首次规则发牌
	BombWeightMap map[int32]int32 //炸弹权重列表
	BombRandRange []int32         //3个玩家随机的权重范围，共6个值

	//非首次发牌规则洗牌
	Seqs    []int32 //不洗牌牌堆
	SizeX   int     //洗牌力度
	CntList []int32 //655的方式发牌
}

type BxpShuffle struct {
	c             *ShuffleConfig
	wList         map[int32]int32 //炸弹权重表
	useMap        map[int32]bool  //牌是否已使用
	userRWeight   []int32         //3个玩家随机的权重
	userRBombList map[int][]int32 //3个玩家随机的炸弹列表
	deck          map[int][]int32 //3个玩家的手牌
}

func NewBxpShuffle(c *ShuffleConfig) *BxpShuffle {
	s := &BxpShuffle{}
	s.c = c
	s.wList = make(map[int32]int32)
	s.useMap = make(map[int32]bool)
	s.userRBombList = make(map[int][]int32)
	s.deck = make(map[int][]int32)
	s.userRWeight = make([]int32, 3)
	return s
}

func (s *BxpShuffle) init() {

	if len(s.c.BombWeightMap) > 0 {
		for i := int32(1); i < _MaxLv; i++ {
			if val, ok := s.c.BombWeightMap[i]; ok {
				s.wList[i] = val
			}
			s.useMap[i] = false
		}
	}

	if arr := s.c.BombRandRange; len(arr) >= 6 {
		for i := 0; i < 3; i++ {
			s.userRWeight[i] = int32(base.RandRange(int(arr[2*i]), int(arr[2*i+1])))
		}
	}

}

func (s *BxpShuffle) DispatchDeckBxp() (a, b, c, bottom []int32) {
	s.init()

	if s.c.UseFirstRule {
		return s.dispatchDeckBxp1()
	} else {
		return s.dispatchDeckBxp2()
	}
}

//洗牌方式1：使用首次发牌规则发牌
func (s *BxpShuffle) dispatchDeckBxp1() (a, b, c, bottom []int32) {
	s.randBombList()
	s.makeDeck()

	if s.checkOk() {
		a, b, c, bottom = RefreshCards(s.deck[0]), RefreshCards(s.deck[1]), RefreshCards(s.deck[2]), s.deck[3]
	} else {
		tmp := base.SliceShuffle(base.SliceCopy(OneDeck()))
		a, b, c, bottom = RefreshCards(tmp[0:17]), RefreshCards(tmp[17:34]), RefreshCards(tmp[34:51]), tmp[51:54]
	}
	return
}

func (s *BxpShuffle) randBombList() {
	for i := 0; i < 3; i++ {
		s.randUserBombList(i)
	}
}

func (s *BxpShuffle) randUserBombList(i int) {
	var (
		weight     = s.userRWeight[i]
		rWeightSum = int32(0)
		rBombList  = []int32(nil)
	)

	//step1:校验是否比最小的炸弹权值还小
	if s.isSmallTheMin(weight) {
		return
	}

	//step2:开始随机炸弹
	for {
		//先随机一个炸弹
		rBombList = s.randBombListByCnt(1)
		rWeightSum = s.calcBombWeightSum(s.userRBombList[i]) + s.calcBombWeightSum(rBombList) //bomb权重和

		if rWeightSum <= weight {
			s.userRBombList[i] = append(s.userRBombList[i], rBombList...)
			continue
		} else if len(rBombList) > 1 && rWeightSum > weight {
			s.cancelUse(rBombList[0])
			rBombList = base.SliceDel(rBombList, rBombList[0])
			rWeightSum = s.calcBombWeightSum(s.userRBombList[i])
			return
		} else {
			return
		}
	}
}

func (s *BxpShuffle) randBombListByCnt(cnt int) (result []int32) {
	randCnt := 0
	for {
		if randCnt >= cnt || s.checkUseOver() {
			return
		}

		if lv := int32(base.RandRange(1, _MaxLv)); !s.useMap[lv] {
			randCnt++
			s.useCard(lv)
			result = append(result, lv)
		}
	}
}

func (s *BxpShuffle) calcBombWeightSum(bombList []int32) int32 {
	sum := int32(0)
	for _, lv := range bombList {
		sum += s.wList[lv]
	}
	return sum
}

func (s *BxpShuffle) checkUseOver() bool {
	for lv, use := range s.useMap {
		if lv != _UnknownLevel && !use {
			return false
		}
	}
	return true
}

func (s *BxpShuffle) isSmallTheMin(weight int32) bool {
	min := int32(0)
	for _, v := range s.wList {
		if v < min && s.useMap[v] == false {
			min = v
		}
	}
	return weight < min
}

func (s *BxpShuffle) useCard(card int32) {
	s.useMap[card] = true
}

func (s *BxpShuffle) cancelUse(card int32) {
	s.useMap[card] = false
}

func (s *BxpShuffle) makeDeck() {
	var (
		hands    = [3][]int32{}
		bottom   = []int32(nil)
		bombSeqs = []int32(nil)
		deck     = base.SliceShuffle(OneDeck())
	)

	for i := 0; i < 3; i++ {
		hands[i] = []int32(nil)
		for _, lv := range s.userRBombList[i] {
			seqs := getBombByLevel(int(lv))
			hands[i] = append(hands[i], seqs...)
			bombSeqs = append(bombSeqs, seqs...)
		}
	}

	if base.BAllInA(deck, bombSeqs) && base.IsUnique(bombSeqs) {

		bottom = base.AExceptB(deck, bombSeqs)
		for i := 0; i < 3; i++ {
			end := len(hands[i])
			seqs := base.SliceNSequence(bottom, 0, 17-end)
			bottom = base.SliceDel(bottom, seqs...)
			hands[i] = append(hands[i], seqs...)

			s.deck[i] = append(s.deck[i], hands[i]...)
		}
		s.deck[3] = append(s.deck[3], bottom...)
	}

}

//校验合法性
func (s *BxpShuffle) checkOk() bool {
	check := []int32(nil)
	for i := 0; i < 4; i++ {
		check = append(check, s.deck[i]...)
	}
	isOneDeck := IsOneDeck(check)
	return isOneDeck
}

// ----------------------------------------------------------------------------------------------------
/*
	不洗牌算法2: 上次打的牌不洗，随机抽取长度为SizeX的连续牌堆，放在牌堆的最上方，再按665的方式发给每个玩家
*/

//洗牌方式2：非首次发牌规则洗牌 (使用洗牌力度按655的方式发牌)
func (s *BxpShuffle) dispatchDeckBxp2() (a, b, c, bottom []int32) {
	var (
		tmp   = []int32(nil)
		hands = map[int][]int32{} //3个玩家手牌
	)

	if IsOneDeck(s.c.Seqs) {
		tmp = base.SliceCopy(s.c.Seqs)
	} else {
		tmp = base.SliceShuffle(OneDeck())
	}

	cards := s.genBxpCards(tmp, s.c.SizeX)

	//分别给3玩家每次发n张牌
	start := 0
	for _, n := range s.c.CntList {
		for i := 0; i < 3; i++ {
			hands[i] = append(hands[i], base.SliceNSequence(cards, start, int(n))...)
			start += int(n)
		}
	}

	a, b, c, bottom = RefreshCards(hands[0]), RefreshCards(hands[1]), RefreshCards(hands[2]), cards[51:54]
	return
}

// 不洗牌生成牌堆
func (s *BxpShuffle) genBxpCards(cards []int32, SizeX int) (result []int32) {
	// 随机抽取长度为SizeX的连续牌堆，放在牌堆的最上方
	deck := base.SliceCopy(cards)
	part := base.SliceRangNSequence(deck, SizeX)
	result = base.SliceReBuild(deck, part)
	return
}
