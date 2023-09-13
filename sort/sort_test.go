package sort

import (
	"sort"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	a := sort.IntSlice{5, 2, 4, 6, 1, 3}
	InsertionSort(a)
	t.Error(a)
}

func TestMerge(t *testing.T) {
	a := []int{8, 4, 2, 5, 7, 9, 3, 2, 6, 1}
	MergeSort(a, 0, len(a)-1)
	t.Error(a)
}
