package dfs

import (
	"testing"
	"time"

	"github.com/go-kratos/kratos/pkg/log"
)

func TestPermute(t *testing.T) {
	var (
		_arr    = []int32(nil)
		_num    = 2  //2个dice
		_maxCnt = 15 //15个chess
	)

	for i := int32(1); i <= int32(_maxCnt); i++ {
		_arr = append(_arr, i)
	}
	start := time.Now()
	ret := Permute(_arr, _num)
	log.Info("use time: %v", time.Since(start).Milliseconds())
	log.Info("hello world, cnt:%+v _num:%+v \n%v", len(ret), _num, ret)
	return
}

func Permute(nums []int32, k int) [][]int32 {
	var res [][]int32
	if len(nums) == 0 {
		return res
	}
	var tmp []int32
	var visited = make([]bool, len(nums))
	backtracking(nums, k, &res, tmp, visited)
	return res
}

func backtracking(nums []int32, k int, res *[][]int32, tmp []int32, visited []bool) {
	// 成功找到一组
	if len(tmp) == k /*len(nums)*/ {
		var c = make([]int32, len(tmp))
		copy(c, tmp)
		*res = append(*res, c)
		return
	}
	// 回溯
	for i := 0; i < len(nums); i++ {
		//if !visited[i] {
		//	visited[i] = true

		//能够移动？
		//if !checkCanMove() {
		//	continue
		//}
		//尝试移动一次
		//moveOne()

		tmp = append(tmp, nums[i])
		backtracking(nums, k, res, tmp, visited)
		tmp = tmp[:len(tmp)-1]

		//回退之前移动的位置
		//backOne()

		//	visited[i] = false
		//}
	}
}
