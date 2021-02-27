package util

import (
	"bytes"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// RandIntn 获取一个 0 ~ n 之间的随机值
func RandIntn(n int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return r.Intn(n)
}

// GetRandString 生成n位随机数字字符串
func GetRandString(n int) string {
	var buffer bytes.Buffer
	for i := 0; i < n; i++ {
		buffer.WriteString(strconv.Itoa(RandIntn(10)))
	}

	return buffer.String()
}

// Rand 生成最大范围内随机数
func RandN(value int) int {
	if value == 0 {
		return 0
	}
	return rand.Intn(value)
}

// RangeNum 生成一个区间范围的随机数 [min,max)
func Rands(min, max int) int {
	if min > max {
		max = min
	}
	if max-min == 0 {
		return max
	}
	randNum := rand.Intn(max - min)
	randNum = randNum + min
	//fmt.Printf("rand is %v\n", randNum)
	return randNum
}

//RangeNum 生成一个区间范围的随机数 [min,max]
func Randsc(min, max int) int {
	if min > max {
		max = min
	}
	if max-min == 0 {
		return max
	}
	randNum := rand.Intn(max - min + 1)
	randNum = randNum + min
	//fmt.Printf("rand is %v\n", randNum)
	return randNum
}

func Rand64() int64 {
	randNum := rand.Int63()
	//	fmt.Printf("rand is %v\n", randNum)
	return randNum
}

// Rand 生成最大范围内随机数
func Randn64(value int64) int64 {
	if value == 0 {
		return 0
	}
	randNum := rand.Int63n(value)
	//	fmt.Printf("rand is %v\n", randNum)
	return randNum
}

// RangeNum 生成一个区间范围的随机数 [min,max)
func Rands64(min, max int64) int64 {
	if min > max {
		max = min
	}
	if max-min == 0 {
		return max
	}
	randNum := rand.Int63n(max - min)
	randNum = randNum + min
	//fmt.Printf("rand is %v\n", randNum)
	return randNum
}

//生成一个随机切片
func RandPerm(n int) []int {
	return rand.Perm(n)
}

func OnStartByPid() {
	rand.Seed(time.Now().Unix() + int64(os.Getpid()))
}

func OnStart() {
	rand.Seed(time.Now().Unix())
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
