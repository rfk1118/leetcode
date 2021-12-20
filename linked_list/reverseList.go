package linked_list

// 206. 反转链表
// https://leetcode-cn.com/problems/reverse-linked-list/
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 这个是返回头的
	root := reverseList(head.Next)
	// 这个是重新处理连接的
	head.Next.Next = head
	// 设置当前head为空的
	head.Next = nil
	return root
}
