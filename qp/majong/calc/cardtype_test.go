package calc

import (
	"fmt"
	"testing"
	"time"

	"github.com/znyh/class/qp/majong/calc/hulib"
)

var (
	errCount = 0
	okCount  = 0
	tested   = map[int]bool{}
)

func printCards(cards []int) {
	for i := 0; i < 9; i++ {
		fmt.Printf("%d ", cards[i])
	}
	fmt.Printf("\n")

	for i := 9; i < 18; i++ {
		fmt.Printf("%d ", cards[i])
	}
	fmt.Printf("\n")

	for i := 18; i < 27; i++ {
		fmt.Printf("%d ", cards[i])
	}
	fmt.Printf("\n")

	for i := 27; i < 34; i++ {
		fmt.Printf("%d ", cards[i])
	}
	fmt.Printf("\n")
}

func checkHu(cards []int, max int) {
	for i := 0; i < max; i++ {
		if cards[i] > 4 {
			return
		}
	}

	num := 0
	for i := 0; i < 9; i++ {
		num = num*10 + cards[i]
	}

	_, ok := tested[num]
	if ok {
		return
	}

	tested[num] = true

	for i := 0; i < max; i++ {
		if !hulib.CheckHu(cards, 34) {
			errCount++
			fmt.Printf("测试失败 i=%d\n", i)
			printCards(cards)
		} else {
			okCount++
		}
	}
}

func genAutoTableSub(cards []int, level int) {
	for i := 0; i < 32; i++ {
		index := -1
		if i <= 17 {
			cards[i] += 3
		} else if i <= 24 {
			index = i - 18
		} else {
			index = i - 16
		}

		if index >= 0 {
			cards[index] += 1
			cards[index+1] += 1
			cards[index+2] += 1
		}

		if level == 4 {
			checkHu(cards, 18)
		} else {
			genAutoTableSub(cards, level+1)
		}

		if i <= 17 {
			cards[i] -= 3
		} else {
			cards[index] -= 1
			cards[index+1] -= 1
			cards[index+2] -= 1
		}
	}
}

func testTwoColor() {
	fmt.Println("testing two colors:")
	cards := []int{
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0,
	}

	for i := 0; i < 18; i++ {
		cards[i] = 2
		fmt.Printf("eye %d\n", i+1)
		genAutoTableSub(cards, 1)
		cards[i] = 0
	}
}

func testTime(count int) {
	cards := []int{
		0, 0, 0, 0, 0, 1, 0, 0, 0,
		1, 1, 0, 0, 0, 0, 1, 0, 0,
		0, 0, 1, 0, 0, 0, 0, 0, 0,
		1, 0, 0, 0, 0, 4, 4,
	}

	printCards(cards)
	start := time.Now().Unix()
	for i := 0; i < count; i++ {
		hulib.CheckHu(cards, 34)
	}
	printCards(cards)
	fmt.Println("count=", count, "use time=", time.Now().Unix()-start)
}

func oneOk() {
	cards := []int{
		0, 0, 0, 2, 1, 2, 4, 3, 2,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0,
	}
	if hulib.CheckHu(cards, 34) == false {
		printCards(cards)
		fmt.Printf("第一次测试失败 ...\n\n")
	} else {
		fmt.Printf("第一次测试成功 ...\n\n")
	}
}

func TestHuType(t *testing.T) {
	time.Sleep(3 * time.Second)

	oneOk()
	start := time.Now().Unix()
	fmt.Println("开始测试...")
	testTime(100000000)
	testTwoColor()
	fmt.Println("测试结束...")
	fmt.Printf("失败次数:%d，成功次数:%d\n", errCount, okCount)
	fmt.Println("测试总次数=", okCount+errCount, "用时", time.Now().Unix()-start)

	fmt.Println("done")
}

func TestSimpleType(t *testing.T) {
	var (
		cards = []int{
			0, 0, 0, 2, 3, 3, 4, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 2,
		}

		weaves = []TagWeaveItem{
			{cbWeaveKind: OpAnGang, cbCenterCard: 0x01,},
			{cbWeaveKind: OpPeng, cbCenterCard: 0x21,},
			//{cbWeaveKind: OpEatLeft, cbCenterCard: 0x016,},
		}

		guiIndex = 33
	)

	set := NewCardsSet(cards, weaves, guiIndex)

	printCards(cards)
	fmt.Printf("set.weaves:%+v\nset.size:%d  set.guiIndex:%d set.guiNum:%d set.handcolors:%+v set.weavecolors:%+v\nset.hu:%t\n\n",
		set.weaves, set.size, set.guiIndex, set.guiNum, set.CalcHandsColor(), set.CalcWeavesColor(), set.hu)

	fmt.Printf("pengpenghu::%t\n", set.CheckPengPengHu())
	fmt.Printf("7dui:%t\n:", set.Check7dui())
	fmt.Printf("HuType:0x%x\n:", HuType(set.elem, weaves, set.guiIndex, MAXINDEX))
}

func TestTing(t *testing.T) {
	cards := []int{
		0, 0, 0, 2, 0, 2, 4, 3, 2,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0,
	}

	printCards(cards)
	set := NewCardsSet(cards, nil, 33)
	fmt.Printf("ting:%+v\n:", set.ting)
}

func TestOutTing(t *testing.T) {
	cards := []int{
		0, 0, 0, 2, 0, 2, 4, 3, 2,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 1,
	}

	printCards(cards)
	set := NewCardsSet(cards, nil, 33)
	fmt.Printf("OutTingTbl:%+v\n:", OutTingTbl(set.elem, set.guiIndex))
}

func TestCardType(t *testing.T) {
	cards := []int{
		0, 0, 0, 1, 2, 3, 3, 3, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 2,
	}

	printCards(cards)
	set := NewCardsSet(cards, nil, 33)
	fmt.Printf("HuType:0x%x\n:", HuType(set.elem, nil, set.guiIndex, MAXINDEX))
}

func TestEat(t *testing.T) {
	cards := []int{
		0, 1, 1, 1, 1, 1, 1, 1, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 2,
	}
	printCards(cards)
	set := NewCardsSet(cards, nil, 33)
	fmt.Printf("ting:0x%x\n:", set.CheckEat(3))

}
