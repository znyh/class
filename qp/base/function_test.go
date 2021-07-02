package base

import (
	"fmt"
	"testing"
)

func TestHexStr2Int32(t *testing.T) {

	src := "  0x01,  0x25, 0x36,  0x47  "
	seqs := HexStr2Int32(src)

	fmt.Printf("src:%s\n", src)

	for _, v := range seqs {
		fmt.Printf("0x%x,", v)
	}
	fmt.Printf("\n\n")

}

func TestCardsRecorder(t *testing.T) {
	hands := map[int32][]int32{}
	others := map[int32][]int32{}
	recorder := map[int32][]int{}

	hands[0] = []int32{1, 2, 3, 4, 5, 6}
	hands[1] = []int32{2, 3, 4, 5, 6, 7}
	hands[2] = []int32{3, 4, 5, 6, 7, 9}

	for i := int32(0); i < 3; i++ {
		for j := int32(0); j < 3; j++ {
			if i != j {
				others[i] = append(others[i], hands[j]...)
			}
		}

		tmp := [16]int{0}
		for _, c := range others[i] {
			number := int(c & 0x0f)
			tmp[number]++
		}

		recorder[i] = append(recorder[i], tmp[1:]...)
	}

	for i := int32(0); i < 3; i++ {
		fmt.Printf("i:%d ==> others:%+v\n", i, others[i])
		fmt.Printf("i:%d ==> record:%+v\n\n", i, recorder[i])
	}
}

func TestDfs(t *testing.T) {
	var a = []int32{11, 12, 13, 14}
	var k = 2
	var dst = Dfs(a, k)
	fmt.Printf("dst:%+v\n", dst)
}

func TestPermute(t *testing.T) {
	nums := []int32{1, 2, 3}
	dst := Permute(nums)
	fmt.Printf("dst:%+v\n", dst)
}

//已知炸弹权重列表candidates和玩家的炸弹权重target，求该玩家能获取到所有的炸弹列表
func TestCombinationSum2(t *testing.T) {
	var (
		candidates = []int{0, 300, 300, 300, 310, 310, 320, 320, 330, 330, 335, 340, 345, 360, 370}
		target     = 650
	)
	dst := combinationSum2(candidates, target)
	fmt.Printf("dst:%+v\n", dst) //得到的炸弹列表 ==》 下标对应牌等级
}

func TestSliceDel(t *testing.T) {
	nums := []int32{1, 2, 3, 4, 5}
	del := []int32{2, 3, 6, 7, 8}
	dst := SliceDel(nums, del...)
	fmt.Printf("dst:%+v\n", dst)
}
