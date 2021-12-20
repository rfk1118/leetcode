package linked_list

// 1721. 交换链表中的节点
// https://leetcode-cn.com/problems/swapping-nodes-in-a-linked-list/
func swapNodes(head *ListNode, k int) *ListNode {
	var x, y *ListNode
	var handler func(node *ListNode, a int) int
	handler = func(node *ListNode, a int) int {
		if nil == node {
			return 0
		}
		a = a - 1
		if a == 0 {
			x = node
		}
		i := handler(node.Next, a)
		if i+1 == k {
			y = node
		}
		return i + 1
	}

	handler(head, k)
	x.Val, y.Val = y.Val, x.Val
	return head
}

func swapNodesV2(head *ListNode, k int) *ListNode {
	var zK, dK *ListNode
	x := head
	for i := 0; i < k; i++ {
		x = x.Next
	}
	zK = x
	y := head
	for ; x != nil; x = x.Next {
		x = x.Next
		y = y.Next
	}
	dK = y
	zK.Val, dK.Val = dK.Val, zK.Val
	return head
}
