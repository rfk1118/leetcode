package top

import (
	"sort"
	"testing"
)

func Test_isMatch(t *testing.T) {
	isMatch("aab", "c*a*b")
	sort.Ints([]int{1, 3, 5, 7, 2, 3, 1})
}
