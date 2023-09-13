package sort

import (
	"math"
	"sort"
)

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

func merge(data []int, p, q, r int) {
	left := make([]int, q-p+2)
	right := make([]int, r-q+1)
	copy(left, data[p:q+1])
	copy(right, data[q+1:r+1])
	left[len(left)-1] = int(math.Inf(1))
	right[len(right)-1] = int(math.Inf(1))

	leftIndex := 0
	rightIndex := 0

	for i := p; i < r+1; i++ {
		if left[leftIndex] < right[rightIndex] {
			data[i] = left[leftIndex]
			leftIndex += 1
		} else {
			data[i] = right[rightIndex]
			rightIndex += 1
		}
	}
}

func MergeSort(data []int, p, r int) {
	if p < r {
		q := (p + r) / 2
		MergeSort(data, p, q)
		MergeSort(data, q+1, r)
		merge(data, p, q, r)
	}
}
