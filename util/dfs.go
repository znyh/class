package util

import (
	"sort"
)

//在a中取m个元素的所有组合
func Dfs(a []int32, k int) [][]int32 {
	if len(a) < k {
		return nil
	}
	var (
		n   = len(a)
		vis = []int32(nil)
		re  = []int32(nil)
		dst = [][]int32(nil)
	)

	for i := 0; i < n; i++ {
		vis = append(vis, 0)
		re = append(re, 0)
	}

	dfs(a, vis, re, &dst, n, k, 0, 0, )

	return dst
}

//参数step代表选取第几个数字，参数start代表从集合的第几个开始选
func dfs(a, vis, re []int32, dst *[][]int32, n, k, start, step int) {
	//如果选够了k个就输出
	if step == k {
		var c = make([]int32, len(re))
		copy(c, re)
		*dst = append(*dst, c[:k])
	}

	for i := start; i < n; i++ {
		if vis[i] == 1 {
			continue
		}

		vis[i] = 1
		re[step] = a[i]

		dfs(a, vis, re, dst, n, k, i+1, step+1, )
		vis[i] = 0
	}

	return
}

//全排列
func Permute(nums []int32) [][]int32 {
	var res [][]int32
	if len(nums) == 0 {
		return res
	}
	var tmp []int32
	var visited = make([]bool, len(nums))
	backtracking(nums, &res, tmp, visited)
	return res
}

func backtracking(nums []int32, res *[][]int32, tmp []int32, visited []bool) {
	// 成功找到一组
	if len(tmp) == len(nums) {
		var c = make([]int32, len(tmp))
		copy(c, tmp)
		*res = append(*res, c)
		return
	}
	// 回溯
	for i := 0; i < len(nums); i++ {
		if !visited[i] {
			visited[i] = true
			tmp = append(tmp, nums[i])
			backtracking(nums, res, tmp, visited)
			tmp = tmp[:len(tmp)-1]
			visited[i] = false
		}
	}
}

//组合总和
func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	if len(candidates) == 0 {
		return res
	}
	sort.Ints(candidates)
	var tmp []int
	var visited = make([]bool, len(candidates))
	//backtracking2(candidates, &res, target, tmp, 0, visited)
	backtracking2(candidates, &res, target, tmp, 1, visited)
	return res
}

func backtracking2(candidates []int, res *[][]int, target int, tmp []int, index int, visited []bool) {
	if target < 0 {
		return
	}
	if target == 0 {
		var c = make([]int, len(tmp))
		copy(c, tmp)
		*res = append(*res, c)
		return
	}
	for i := index; i < len(candidates); i++ {
		if i > 0 && candidates[i] == candidates[i-1] && !visited[i-1] {
			continue
		}
		visited[i] = true
		//tmp = append(tmp, candidates[i])  //记录元素值
		tmp = append(tmp, i) //记录元素的下标
		backtracking2(candidates, res, target-candidates[i], tmp, i+1, visited)
		tmp = tmp[:len(tmp)-1]
		visited[i] = false
	}
}
