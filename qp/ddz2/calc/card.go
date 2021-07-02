package calc

import (
	"math/rand"
	"sort"
)

const (
	MASKCOLOR = 0xF0
	MASKVALUE = 0x0F
)

const (
	MAXCARD  = 16
	MAXCOLOR = 4
)

var (
	oneCards = []int32{
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, //方块
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, //梅花
		0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, //红桃
		0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, //黑桃
		0x4e, 0x4f,
	}

	descCards = map[int32]string{
		0x01: "方块A", 0x02: "方块2", 0x03: "方块3", 0x04: "方块4", 0x05: "方块5", 0x06: "方块6", 0x07: "方块7", 0x08: "方块8", 0x09: "方块9", 0x0a: "方块10", 0x0b: "方块J", 0x0c: "方块Q", 0x0d: "方块K",
		0x11: "梅花A", 0x12: "梅花2", 0x13: "梅花3", 0x14: "梅花4", 0x15: "梅花5", 0x16: "梅花6", 0x17: "梅花7", 0x18: "梅花8", 0x19: "梅花9", 0x1a: "梅花10", 0x1b: "梅花J", 0x1c: "梅花Q", 0x1d: "梅花K",
		0x21: "红桃A", 0x22: "红桃2", 0x23: "红桃3", 0x24: "红桃4", 0x25: "红桃5", 0x26: "红桃6", 0x27: "红桃7", 0x28: "红桃8", 0x29: "红桃9", 0x2a: "红桃10", 0x2b: "红桃J", 0x2c: "红桃Q", 0x2d: "红桃K",
		0x31: "黑桃A", 0x32: "黑桃2", 0x33: "黑桃3", 0x34: "黑桃4", 0x35: "黑桃5", 0x36: "黑桃6", 0x37: "黑桃7", 0x38: "黑桃8", 0x39: "黑桃9", 0x3a: "黑桃10", 0x3b: "黑桃J", 0x3c: "黑桃Q", 0x3d: "黑桃K",
		0x4e: "小王", 0x4f: "大王",
	}
)

func OneDeck() []int32 {
	deck := make([]int32, len(oneCards))
	copy(deck, oneCards)
	return deck
}

func ShuffleDeck() []int32 {
	deck := OneDeck()
	for i := len(deck) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
	return deck
}

func SortCardsData(cbCardsData [] int32) []int32 {
	sort.Slice(cbCardsData, func(i, j int) bool {
		return cbCardsData[i] < cbCardsData[j]
	})
	return cbCardsData
}

func ValidCardsData(cbCardsData []int32) bool {
	for _, cbCardData := range cbCardsData {
		if _, ok := descCards[cbCardData]; !ok {
			return false
		}
	}
	return true
}

func ValidCardData(cbCardData int32) bool {
	cbValue := cbCardData & MASKVALUE
	cbColor := (cbCardData & MASKCOLOR) >> 4
	return ((cbValue >= 1) && (cbValue <= 0x0d) && (cbColor <= 3)) ||
		((cbValue >= 0x0e) && (cbValue <= 0x0f) && (cbColor == 4))
}

func ToIndex(cbCardsData []int32) []int {
	array := make([]int, MAXCARD)
	for _, cbCardData := range cbCardsData {
		if index := ToCardIndex(cbCardData); index >= 0 && index < MAXCARD {
			array[index]++
		}
	}
	return array
}

func ToCardIndex(cbCardData int32) int {
	cbValue := int(cbCardData & MASKVALUE)
	if cbValue >= 0x03 && cbValue <= 0x0d {
		return cbValue - 3
	} else if cbValue >= 0x01 && cbValue <= 0x02 {
		return cbValue + 10
	} else if cbValue >= 0x0e && cbValue <= 0x0f {
		return cbValue - 1
	}
	return -1
}

// `A` => 11
func CardAIdx() int {
	return ToCardIndex(0x01)
}

// `2` => 12
func Card2Idx() int {
	return ToCardIndex(0x02)
}

// `joker` => 14
func CardJokerIdx() int {
	return ToCardIndex(0x4f)
}
