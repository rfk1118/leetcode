package linked_list

import (
	"testing"
)

func Test_oddEvenList(t *testing.T) {
	// [1,2,3,4,5]
	list := generateList([]int{1, 2, 3, 4, 5})
	oddEvenList(list)
}
