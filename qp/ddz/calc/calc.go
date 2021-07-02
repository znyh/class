package calc

import (
	"fmt"
	"sort"
	"strings"

	"github.com/znyh/class/qp/base"
)

var (
	containBombLimit = false //2人亲友房炸弹不可以组成飞机和4代2
)

func Init(isCustomArena bool) {
	containBombLimit = isCustomArena
}

//经典场发牌
func DispatchDeck() [][]int32 {
	tmp := base.SliceCopy(oneCards)
	tmp = base.SliceShuffle(tmp)

	a, b, c, bottom := RefreshCards(tmp[0:17]), RefreshCards(tmp[17:34]), RefreshCards(tmp[34:51]), tmp[51:54]

	return [][]int32{a, b, c, bottom}
}

//不洗牌场发牌 seqs []int32, useFirstRule bool, bombWeightMap map[int32]int32
func DispatchDeckBxp(cfg *ShuffleConfig) [][]int32 {
	s := NewBxpShuffle(cfg)
	a, b, c, bottom := s.DispatchDeckBxp()
	return [][]int32{a, b, c, bottom}
}

//2人斗地主亲友场发牌
func DispatchDeck2RenCustom() [][]int32 {
	var (
		deck = []int32(nil)
		del  = []int32{0x03, 0x13, 0x23, 0x33, 0x04, 0x14, 0x24, 0x34,}
	)
	deck = base.SliceDel(OneDeck(), del...)
	deck = base.SliceShuffle(deck)

	a, b, bottom := RefreshCards(deck[0:17]), RefreshCards(deck[17:34]), deck[34:37]
	return [][]int32{a, b, bottom}
}

//牌型比较大小, 返回true 则 牌型A > 牌型B
func ABiggerB(a, b []int32) bool {
	sa := NewCardsSet(a)
	sb := NewCardsSet(b)

	return sa.IsBigger(sb)
}

//牌型比较大小, 返回true 则 牌型A > 牌型B
func HaveBigger(a, b []int32) bool {
	sa := NewCardsSet(a)
	sb := NewCardsSet(b)

	return sa.HaveBigger(sb)
}

// 计算牌型
func GetType(seqs []int32) int {
	return NewCardsSet(seqs).ct
}

// 按牌值和花色降序排序  花色大小：黑桃-红桃-梅花-方块
func RefreshCards(seqs []int32) (result []int32) {
	if len(seqs) <= 1 {
		return seqs
	}

	m := map[int][]Card{}
	set := NewCardsSet(seqs)

	sort.Slice(set.cards, func(i, j int) bool {
		return set.cards[i].level > set.cards[j].level
	})

	for _, card := range set.cards {
		lv := card.level
		m[lv] = append(m[lv], card)
	}

	for _, cards := range m {
		sort.Slice(cards, func(i, j int) bool {
			return cards[i].color > cards[j].color
		})
	}

	for lv := 20; lv >= 0; lv-- {
		if cards, ok := m[lv]; ok {
			for _, c := range cards {
				result = append(result, c.Seq)
			}
		}
	}

	return
}

func CardsDesc(seqs []int32) string {
	b := strings.Builder{}
	for k, v := range seqs {
		if k == 0 {
			b.WriteString(cardDesc(v))
		} else {
			b.WriteString("," + cardDesc(v))
		}
	}
	return b.String()
}

func cardDesc(seq int32) string {
	if desc, ok := descCards[seq]; ok {
		return desc
	}
	return "$"
}

func CardsDescHex(seqs []int32) string {
	b := strings.Builder{}
	for k, v := range seqs {
		if k == 0 {
			b.WriteString(cardDescHex(v))
		} else {
			b.WriteString("," + cardDescHex(v))
		}
	}
	return b.String()
}

func cardDescHex(seq int32) string {
	if _, ok := descCards[seq]; ok {
		return fmt.Sprintf("0x%x", seq)
	}
	return ""
}

func CardsSetDesc(seqs []int32) string {
	set := NewCardsSet(seqs)
	return fmt.Sprintf("%+v", set)
}

// 找出所有连续长度为length且每个元素有count个的组合
func FindContinue(seqs []int32, length, count int) [][]int32 {
	if count <= 0 || length <= 0 {
		return nil
	}
	cardsRet := [][]int32(nil)
	set := NewCardsSet(seqs)

	//所有的单张
	if count == 1 && length == 1 {
		for lv, c := range set.count {
			if c > 0 {
				cards, _ := set.getSeqsByLevel(lv, count)
				cardsRet = append(cardsRet, cards)
			}
		}
		return cardsRet
	}

	//所有的连续组合
	i, j, continues := 1, 1, 0
	for i = 1; i < _MaxLv-2; i++ {
		if set.count[i] < count {
			continue
		}
		continues = 0
		for j = i; j < _MaxLv-2 && continues <= length && set.count[j] >= count; j++ {
			continues++
			if continues == length {
				for start := j - continues + 1; start <= j; start++ {
					cards, _ := set.getSeqsByLevel(start, count)
					cardsRet = append(cardsRet, cards)
				}
				break
			}
		}
	}
	return cardsRet
}
