package linked_list

import (
	"fmt"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	list := generateList([]int{3, 5})
	between := reverseBetween(list, 1, 1)
	fmt.Println(between)
}
