package linked_list

// 86. 分隔链表
// https://leetcode-cn.com/problems/partition-list/
func partition(head *ListNode, x int) *ListNode {
	if nil == head {
		return head
	}
	dumpV1 := &ListNode{}
	dumpV2 := &ListNode{}
	cV1 := dumpV1
	cV2 := dumpV2
	for head != nil {
		if head.Val < x {
			cV1.Next = &ListNode{Val: head.Val}
			cV1 = cV1.Next
		} else {
			cV2.Next = &ListNode{Val: head.Val}
			cV2 = cV2.Next
		}
		head = head.Next
	}
	cV1.Next = dumpV2.Next
	return dumpV1.Next
}
