package base

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net"
	"sort"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

//GetTick 毫秒级时间戳
func GetTick() int64 {
	return time.Now().UnixNano() / 1e6
}

//获取当前时间
func GetCurSec() int64 {
	return time.Now().Unix()
}

//Rand [0, max)
func Rand(max int) int {
	if max <= 0 {
		return 0
	}
	return rand.Intn(max)
}

//RandRange [min, max)
func RandRange(min int, max int) int {
	if min < 0 || max <= 0 {
		return 0
	}
	if min >= max {
		return min
	}
	return min + rand.Intn(max-min)
}

// 在[0,total)里，随机一个连续的长度为size的区间
func GetRandList(size, total int) (result []int32) {
	if total <= 0 {
		return nil
	}

	x := RandRange(0, total)
	for i := 0; i < size; i++ {
		val := (x + i) % total
		result = append(result, int32(val))
	}
	return
}

func IsHit(max int) bool {
	return Rand(100) < max
}

//ToJSON json string
func ToJSON(v interface{}) string {
	j, err := json.Marshal(v)
	if err != nil {
		return err.Error()
	}
	return string(j)
}

//string转到int
func StrToInt(src string) int {
	//log.Printf("StrToInt:%s.", src)
	dst, err := strconv.Atoi(src)
	if err != nil {
		log.Printf("str to int error(%v).", err)
		return 0
	}
	return dst
}

//string转到int64
func StrToInt64(src string) int64 {
	//	log.Print("src:%s.", src)
	dst, err := strconv.ParseInt(src, 10, 64)
	if err != nil {
		log.Printf("str to int64 error(%v).", err)
		return 0
	}
	return dst
}

//int转到string
func IntToStr(src int) string {
	return strconv.Itoa(src)
}

//int64转到string
func Int64ToStr(src int64) string {
	return strconv.FormatInt(src, 10)
}

//string转到int32
func StrToInt32(src string) int32 {
	dst, err := strconv.Atoi(src)
	if err != nil {
		log.Printf("str to int32 error(%v).", err)
		return 0
	}
	return int32(dst)
}

//int转到string
func Int32ToStr(src int32) string {
	return strconv.Itoa(int(src))
}

//ToString .
func ToString(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

// EarthDistance 计算两个经纬度之间的距离, 单位为米
// 参数列表：纬度1、经度1、纬度2、经度2
func EarthDistance(lat1, lng1, lat2, lng2 float64) float64 {
	radius := float64(6371000) // 赤道半径:6378137
	rad := math.Pi / 180.0

	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad

	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))

	return dist * radius
}

func SliceShuffle(slice []int32) []int32 {
	for i := len(slice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func SliceContain(slice []int32, values ...int32) bool {
	if len(slice) == 0 || len(values) == 0 {
		return false
	}

	for _, val := range values {
		if !sliceContain(slice, val) {
			return false
		}
	}

	return true
}

func sliceContain(slice []int32, value int32) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func SliceCopy(slice []int32) []int32 {
	dst := make([]int32, len(slice))
	copy(dst, slice)
	return dst
}

func SliceDel(slice []int32, values ...int32) []int32 {
	if len(slice) == 0 || len(values) == 0 {
		return slice
	}

	for _, val := range values {
		slice = sliceDel(slice, val)
	}
	return slice
}

func sliceDel(slice []int32, val int32) []int32 {
	for k, v := range slice {
		if v == val {
			return append(append([]int32{}, slice[:k]...), slice[k+1:]...)
		}
	}
	return slice
}

func SliceUnique(slice []int32) []int32 {
	var (
		unique = []int32(nil)
		m      = make(map[int32]bool)
	)
	for _, v := range slice {
		if _, exist := m[v]; !exist {
			m[v] = true
			unique = append(unique, v)
		}
	}
	return unique
}

//升序
func SliceSortRise(slice []int32) []int32 {
	sort.SliceStable(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	return slice
}

//降序
func SliceSortDesc(slice []int32) []int32 {
	sort.SliceStable(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})
	return slice
}

//两个切片排序后的元素是否全都一样
func SliceSortEqual(slice1, slice2 []int32) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	s1 := SliceSortDesc(SliceCopy(slice1))
	s2 := SliceSortDesc(SliceCopy(slice2))
	for k, v := range s1 {
		if s2[k] != v {
			return false
		}
	}
	return true
}

func SliceReverse(slice []int32) []int32 {
	size := len(slice)
	for i := 0; i < size/2; i++ {
		slice[i], slice[size-i-1] = slice[size-i-1], slice[i]
	}
	return slice
}

//取slice[start,start+n]上的元素
func SliceNSequence(slice []int32, start, n int) []int32 {
	if start+n >= len(slice) {
		return make([]int32, n)
	}
	return slice[start : start+n]
}

// 随机取长度为n的连续序列
func SliceRangNSequence(slice []int32, n int) (result []int32) {
	if n > len(slice) {
		return make([]int32, n)
	}
	list := GetRandList(n, len(slice))
	for _, v := range list {
		result = append(result, slice[v])
	}
	return result
}

// 抽取slice中的part部分，将part部分放到slice的最前面
func SliceReBuild(slice, part []int32) (result []int32) {

	if !BAllInA(slice, part) {
		return
	}

	left := AExceptB(slice, part)
	result = append(result, part...)
	result = append(result, left...)

	return
}

func MinFloat64(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func MInInt32(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

func MaxInt32(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func GetGameNo(roomid, tableid int32) int64 {
	t := GetTick() / 1000
	return t<<16 + int64(roomid)<<8 + int64(tableid)
}

//Copy deep copy struct 注意：只对结构体大写变量生效!
func Copy(src, des interface{}) error {
	j, err := json.Marshal(src)
	if err != nil {
		return err
	}

	return json.Unmarshal(j, des)
}

func AExceptB(a, b []int32) []int32 {
	ret := make([]int32, 0)
	for _, v := range a {
		if !isIn(v, b) {
			ret = append(ret, v)
		}
	}

	return ret
}

func BAllInA(a, b []int32) bool {
	for _, v := range b {
		if !isIn(v, a) {
			return false
		}
	}

	return true
}

//数组值都是唯一的
func IsUnique(a []int32) bool {
	m := map[int32]bool{}
	for _, v := range a {
		if m[v] {
			return false
		}

		m[v] = true
	}

	return true
}

func isIn(e int32, ems []int32) bool {
	for _, v := range ems {
		if e == v {
			return true
		}
	}
	return false
}

func IsValidIdx(slice []int32, idx []int32) bool {
	size := len(slice)
	for _, id := range idx {
		if id >= int32(size) || id < 0 {
			return false
		}
	}
	return true
}

func GetElemByIdx(slice []int32, idx []int32) []int32 {
	size := len(slice)
	ret := []int32(nil)

	for _, id := range idx {
		if id >= int32(size) || id < 0 {
			return nil
		}
		ret = append(ret, slice[id])
	}
	return ret
}

//指定概率的随机事件 返回落到probability对应的概率的下标
func RandEvent(probability []int32) int {
	r := rand.Int31n(99)
	start, end := int32(0), int32(0)

	for k, v := range probability {
		end = start + v - 1
		if r >= start && r <= end {
			return k

		}
		start += v
	}
	return 0
}

//字符串表示的16进制转化为整形
func HexStr2Int32(src string) (dst []int32) {
	if len(src) <= 0 {
		return
	}

	tmp := strings.TrimSpace(src)
	s := strings.Split(tmp, ",")

	for _, v := range s {
		vv := strings.TrimSpace(v)
		if val, err := strconv.ParseInt(vv, 0, 0); err == nil {
			dst = append(dst, int32(val))
		} else {
			return []int32{}
		}
	}

	return dst
}

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

func GetIP() string {

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "127.0.0.1"
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return "127.0.0.1"
}
