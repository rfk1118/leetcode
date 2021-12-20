package linked_list

// https://leetcode-cn.com/problems/merge-two-sorted-lists/
// 21. 合并两个有序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if nil == l1 {
		return l2
	}
	if nil == l2 {
		return l1
	}
	root := &ListNode{}
	cRootIndex := root
	currentL1 := l1
	currentL2 := l2
	for currentL1 != nil && currentL2 != nil {
		if currentL1.Val < currentL2.Val {
			cRootIndex.Next = &ListNode{Val: currentL1.Val}
			currentL1 = currentL1.Next
		} else {
			cRootIndex.Next = &ListNode{Val: currentL2.Val}
			currentL2 = currentL2.Next
		}
		cRootIndex = cRootIndex.Next
	}
	if currentL1 == nil {
		cRootIndex.Next = currentL2
	} else {
		cRootIndex.Next = currentL1
	}
	return root.Next
}
