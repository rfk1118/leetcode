package top

import (
	"fmt"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	var a []int = []int{1, 2, 3}
	newCollection := make([]int, 0)
	newCollection = append(newCollection, a...)
	for _, v := range newCollection {
		fmt.Print(v)
	}
}
