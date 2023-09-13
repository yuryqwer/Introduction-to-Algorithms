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
