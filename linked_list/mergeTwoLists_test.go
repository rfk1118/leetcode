package linked_list

import (
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	//[1,2,4]
	//[1,3,4]
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	_ = mergeTwoLists(l1, l2)
}
