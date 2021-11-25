package linked_list

// 2. 两数相加
// https://leetcode-cn.com/problems/add-two-numbers/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var handler func(n1 *ListNode, n2 *ListNode, add int) *ListNode
	handler = func(n1 *ListNode, n2 *ListNode, add int) *ListNode {
		if n1 == nil && n2 == nil && add == 0 {
			return nil
		}
		node := &ListNode{Val: 0}
		var n1n, n2n *ListNode
		if n1 == nil && n2 == nil {
			node.Val = node.Val + add
			n1n = nil
			n2n = nil
		} else if n1 == nil {
			node.Val = node.Val + n2.Val + add
			n1n = nil
			n2n = n2.Next
		} else if n2 == nil {
			node.Val = node.Val + n1.Val + add
			n2n = nil
			n1n = n1.Next
		} else {
			node.Val = node.Val + n1.Val + n2.Val + add
			n1n = n1.Next
			n2n = n2.Next
		}
		if node.Val >= 10 {
			node.Val = node.Val % 10
			// 9+9+1
			node.Next = handler(n1n, n2n, 1)
		} else {
			node.Next = handler(n1n, n2n, 0)
		}
		return node
	}
	return handler(l1, l2, 0)
}
