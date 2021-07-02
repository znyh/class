package calc

import (
	"fmt"
	"testing"
	"time"
)

func TestGenBxpCards(t *testing.T) {

	const (
		testCnt = 100
	)

	var (
		okCnt  = 0
		errCnt = 0
		start  = time.Now()
	)

	c := &ShuffleConfig{
		UseFirstRule:  true,
		BombWeightMap: make(map[int32]int32),
		BombRandRange: nil,
	}

	WtList := [_MaxLv]int32{0, 300, 300, 300, 310, 310, 320, 320, 330, 330, 335, 340, 345, 360, 370}
	for i := 0; i < _MaxLv; i++ {
		c.BombWeightMap[int32(i)] = WtList[i]
	}
	c.BombRandRange = []int32{300, 480, 300, 480, 300, 1000}

	for i := 0; i < testCnt; i++ {

		s := NewBxpShuffle(c)
		s.init()
		s.dispatchDeckBxp1()

		if s.checkOk() {
			okCnt++
		} else {
			errCnt++
		}
	}

	fmt.Printf("\nerrCnt:%d testCnt:%d useTime:%v\n\n\n", errCnt, testCnt, time.Since(start))

}

func TestAIRecord(t *testing.T) {
	dst := ""

	record := [][]int32{{}, {}, {0x01, 0x02, 0x13, 0x24, 0x35,}, {}, {0x05}}

	str := []string(nil)
	for _, j := range record {
		if len(j) == 0 {
			str = append(str, "PASS")
		} else {
			str = append(str, CardsDescHex(j))
		}
	}

	for k, v := range str {
		if k == 0 {
			dst = v
		} else {
			dst += fmt.Sprintf("\\|%s", v)
		}
	}

	fmt.Printf("dst:%s\n", dst)
}

func TestCardType(t *testing.T) {

	//seqs := []int32{49, 33, 17, 1, 61, 45, 29, 13, 60, 44, 28, 12, 59, 43, 27, 11}
	seqs := []int32{0x03, 0x13, 0x23, 0x04, 0x14, 0x24, 0x05, 0x15, 0x25, 0x06, 0x16, 0x07}

	set := NewCardsSet(seqs)

	fmt.Printf("ty:%d lv:%d length:%d\n", set.ct, set.level, set.length)
}

func TestIsHaveBigger(t *testing.T) {

	//seqs := []int32{0x03, 0x13, 0x23, 0x04, 0x14, 0x24, 0x05, 0x15, 0x25, 0x06, 0x16, 0x07}
	//hand := []int32{0x0a, 0x1a, 0x2a, 0x3a, 0x0b, 0x1b, 0x2b, 0x3b, 0x0c, 0x1c, 0x2c, 0x3c, 0x3d, 0x4e, 0x28}
	//

	//seqs := []int32{78,22,6,37,21,5,36,20,4,35,19,3}
	//hand := []int32{7, 23, 39, 55, 8, 24, 40, 56, 9, 25, 41, 57, 10, 26, 42, 58, 79}

	//seqs := []int32{0x03, 0x13, 0x23, 0x04, 0x14, 0x24, 0x05, 0x15, 0x25, 0x06, 0x17, 0x08}
	//hand := []int32{0x07, 0x17, 0x27, 0x08, 0x18, 0x28, 0x09, 0x19, 0x29, 0x0a, 0x1a, 0x2a}

	seqs := []int32{0x03, 0x13, 0x23, 0x33, 0x14, 0x24, 0x04, 0x34}
	hand := []int32{0x07, 0x17, 0x27, 0x37, 0x18, 0x28, 0x38, 0x08}

	s1 := NewCardsSet(seqs)
	s2 := NewCardsSet(hand)

	isbig := s2.HaveBigger(s1)

	fmt.Printf("s1=> ty:%d lv:%d length:%d\n", s1.ct, s1.level, s1.length)
	fmt.Printf("s2=> ty:%d lv:%d length:%d\n", s2.ct, s2.level, s2.length)
	fmt.Printf("s2 isbig s1 ? ==> %t\n", isbig)
}
