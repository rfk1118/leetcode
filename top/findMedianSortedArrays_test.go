package top

import (
	"fmt"
	"testing"
)

func Test_findMedianSortedArrays(t *testing.T) {
	f := findMedianSortedArrays([]int{1, 2}, []int{1, 2, 3, 4, 5})
	fmt.Print(f)
}
