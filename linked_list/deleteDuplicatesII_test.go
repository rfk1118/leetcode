package linked_list

import (
	"fmt"
	"testing"
)

func Test_deleteDuplicatesII(t *testing.T) {
	var r = []int{1, 2, 3, 3, 4, 4, 5}
	list := generateList(r)
	ii := deleteDuplicatesII(list)
	fmt.Println(ii)
}
