package linked_list

import "math"

// 61. 旋转链表
// https://leetcode-cn.com/problems/rotate-list/
var n int
var max = math.MaxInt64

func rotateRight(head *ListNode, k int) *ListNode {
	// 数量，头，尾
	var rotate func(node *ListNode) (int, *ListNode, *ListNode)
	rotate = func(node *ListNode) (int, *ListNode, *ListNode) {
		n++
		if node.Next == nil {
			max = k % n
			return 1, nil, node
		}
		i, h, t := rotate(node.Next)
		if nil != h {
			return i, h, t
		}
		if i == max {
			h = node.Next
			node.Next = nil
			t.Next = head
		}
		return i + 1, h, t
	}
	_, h, _ := rotate(head)
	return h
}
