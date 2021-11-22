package linked_list

import (
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	// [1,1,2]
	node := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}
	deleteDuplicates(node)

}
