package sort

import "sort"

func InsertionSort(data sort.Interface) {
	n := data.Len()
	if n <= 1 {
		return
	}

	for j := 1; j < n; j++ {
		i := j - 1
		for i >= 0 && data.Less(j, i) {
			data.Swap(i, j)
			i -= 1
			j -= 1
		}
	}
}
