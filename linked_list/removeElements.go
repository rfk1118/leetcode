package linked_list

// https://leetcode-cn.com/problems/remove-linked-list-elements/
// 203. 移除链表元素
func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	if head == nil {
		return nil
	}
	current := head
	for current != nil {
		if current.Next != nil && current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}
