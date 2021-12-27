package top

// 19. 删除链表的倒数第 N 个结点
// https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var handler func(node *ListNode) (int, *ListNode)
	handler = func(node *ListNode) (int, *ListNode) {
		if nil == node {
			return 0, node
		}
		i, ln := handler(node.Next)
		if i+1 == n {
			return i + 1, ln
		}
		node.Next = ln
		return i + 1, node
	}
	_, root := handler(head)
	return root
}

// quick slow
// 快慢步没找到精髓
func removeNthFromEndV2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := head, dummy
	// 先走n
	for x := 0; first != nil && x < n; first = first.Next {
		x++
	}
	for ; first != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}
