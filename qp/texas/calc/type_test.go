package calc

import (
	"fmt"
	"testing"
)

//大小顺序 2 3 4 5 6 7 8 9 10 J(11) Q(12) K(13) A(14)
func TestRefresh(t *testing.T) {
	seqs := []int32{
		01, 02, 03, 04, 05, //方块
		14, 15, //梅花
	}
	set := NewCardSet(seqs)
	fmt.Printf("%+v\n", set)

	for i := 0; i < len(set.cards); i++ {
		fmt.Printf("%d -> %d\n", set.cards[i].level, set.cards[i].color)
	}
}

func TestFlushChain(t *testing.T) {
	seqs := []int32{
		01, 02, 03, 04, 05, 06, 07, //方块
	}
	set := NewCardSet(seqs)
	fmt.Printf("%+v\n", set)

	for i := 0; i < len(set.cards); i++ {
		fmt.Printf("%d -> %d\n", set.cards[i].level, set.cards[i].color)
	}
}

func TestFlush(t *testing.T) {
	seqs := []int32{
		01, 02, 03, 04, 05 + 13, 06, 07, //方块
	}
	set := NewCardSet(seqs)
	fmt.Printf("%+v\n", set)

	for i := 0; i < len(set.cards); i++ {
		fmt.Printf("%d -> %d\n", set.cards[i].level, set.cards[i].color)
	}
}

func TestChain(t *testing.T) {
	seqs := []int32{
		01, 02, 03, 04, 05 + 13, 06 + 13, 07 + 13, //方块
	}
	set := NewCardSet(seqs)
	fmt.Printf("%+v\n", set)

	for i := 0; i < len(set.cards); i++ {
		fmt.Printf("%d -> %d\n", set.cards[i].level, set.cards[i].color)
	}
}
