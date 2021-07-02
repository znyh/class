package calc

import (
	"math/rand"
)

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

func Shuffle(seqs []int32) []int32 {
	dst := make([]int32, len(seqs))
	copy(dst, seqs)

	for i := 0; i < len(seqs); i++ {
		idx := RandRange(i, len(seqs)-i)
		dst[i], dst[idx] = dst[idx], dst[i]
	}
	return dst
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
