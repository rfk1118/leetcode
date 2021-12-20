package linked_list

import (
	"testing"
)

// [9,9,9,9,9,9,9]
// [9,9,9,9]
func Test_addTwoNumbers(t *testing.T) {
	l1 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}}
	l2 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}
	addTwoNumbers(l1, l2)
}
