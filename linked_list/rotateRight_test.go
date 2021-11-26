package linked_list

import (
	"fmt"
	"testing"
)

func Test_rotateRight(t *testing.T) {
	node := &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}
	right := rotateRight(node, 4)
	fmt.Println(right)
}
