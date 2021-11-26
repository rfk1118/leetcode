package linked_list

//  分组？我的第一想法，一会做处理
func reverseKGroup(head *ListNode, k int) *ListNode {
	var reverse func(*ListNode, *ListNode) *ListNode
	reverse = func(a *ListNode, b *ListNode) *ListNode {
		if a.Next == b {
			return a
		}
		node := reverse(a.Next, b)
		a.Next.Next = a
		a.Next = b
		return node
	}
	var reverseKGroup func(node *ListNode, k int) *ListNode
	reverseKGroup = func(node *ListNode, k int) *ListNode {
		if nil == node {
			return node
		}
		// 区间 [a, b) 包含 k 个待反转元素
		a, b := node, node
		for i := 0; i < k; i++ {
			if nil == b {
				return node
			}
			b = b.Next
		}
		// [a,b] -> [b,a]
		newHead := reverse(a, b)
		a.Next = reverseKGroup(b, k)
		return newHead
	}
	return reverseKGroup(head, k)
}
