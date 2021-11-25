package linked_list

// https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
// 19. 删除链表的倒数第 N 个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var remove func(*ListNode) (int, bool)
	remove = func(node *ListNode) (int, bool) {
		if node == nil {
			return 0, false
		}
		i, b := remove(node.Next)
		if b {
			return i, b
		}
		if !b && i == n {
			node.Next = node.Next.Next
			return i + 1, true
		}
		return i + 1, b
	}
	i, b := remove(head)
	if !b && i == n {
		return head.Next
	}
	return head
}
