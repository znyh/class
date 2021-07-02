package calc

import (
	"math/rand"
)

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
