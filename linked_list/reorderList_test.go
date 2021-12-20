package linked_list

import "testing"

func Test_reorderListV2(t *testing.T) {
	list := generateList([]int{1, 2, 3, 4})
	reorderListV2(list)
}
