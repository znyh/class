package util

import (
	"sort"
)

//升序
func Int32SortAscend(slice []int32) []int32 {
	sort.SliceStable(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	return slice
}

//降序
func Int32SortDescend(slice []int32) []int32 {
	sort.SliceStable(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})
	return slice
}

func selectSort(arr []int) {

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i - 1
		tmp := arr[i]
		for j = i - 1; j >= 0; j-- {
			if arr[j] > tmp {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = tmp
	}
}

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	i, j := left, right
	tmp := arr[i]

	for i < j {
		for i < j && arr[j] >= tmp {
			j--
		}
		arr[i], arr[j] = arr[j], arr[i]
		for i < j && arr[i] <= tmp {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[i] = tmp
	quickSort(arr, left, i-1)
	quickSort(arr, i+1, right)
}
