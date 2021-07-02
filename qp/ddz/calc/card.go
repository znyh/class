package calc

import (
	"fmt"
)

/*
	一张牌定义
*/

// Card 一张牌
type Card struct {
	Seq    int32
	color  int
	number int
	level  int
}

func (c *Card) Desc() string {
	return descCards[c.Seq]
}

func (c *Card) calc() {
	if _, find := descCards[c.Seq]; find {
		c.color = cardColor(c.Seq)
		c.number = cardNumber(c.Seq)
		c.level = cardLevel(c.Seq)
	} else {
		fmt.Printf("\r\n\n==> bad seq 0x%0x\n", c.Seq)
	}
}

func cardColor(seq int32) int {
	return int((seq & 0xf0) >> 4)
}

func cardNumber(seq int32) int {
	return int(seq & 0x0f)
}

func cardLevel(seq int32) int {
	return number2Level(cardNumber(seq))
}

func number2Level(n int) int {
	if n >= 3 && n <= 13 {
		return n - 2
	} else if n == 1 || n == 2 {
		return n + 11
	} else if n == 14 || n == 15 {
		return n
	}
	return -1
}

func level2Number(l int) int {
	if l >= 1 && l <= 11 {
		return l + 2
	} else if l == 12 || l == 13 {
		return l - 11
	} else if l == 14 || l == 15 {
		return l
	}
	return -1
}

func getBombByLevel(lv int) (result []int32) {
	if lv <= 0 || lv >= 16 {
		return
	}

	if lv == 14 || lv == 15 {
		return []int32{0x4e, 0x4f}
	}

	number := int32(level2Number(lv))
	return []int32{number, number + 0x10, number + 0x20, number + 0x30}
}

func GetCardLevel(seq int32) int {
	level := cardLevel(seq)
	return level
}

func CompareLevel(seq1 int32, seq2 int32) bool {
	level0 := cardLevel(seq1)
	level1 := cardLevel(seq2)
	return level0 > level1
}
