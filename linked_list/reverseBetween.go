package linked_list

// https://leetcode-cn.com/problems/reverse-linked-list-ii/
// 92. 反转链表 II
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dump := &ListNode{Next: head}
	c := dump
	cN := 0
	var leftNodePre, leftNode, rightNode *ListNode
	for c != nil {
		if nil != leftNode && nil != rightNode {
			break
		}
		if cN == left-1 {
			leftNodePre = c
		}
		if cN == left {
			leftNode = c
		}
		if cN == right {
			rightNode = c.Next
		}
		cN++
		c = c.Next
	}
	var reverseB func(a, b *ListNode) (*ListNode, *ListNode)
	reverseB = func(a, b *ListNode) (*ListNode, *ListNode) {
		if a.Next == b {
			return a, a.Next
		}
		h, t := reverseB(a.Next, b)
		a.Next.Next = a
		a.Next = t
		return h, t
	}
	b, _ := reverseB(leftNode, rightNode)
	leftNodePre.Next = b
	return dump.Next
}

func reverseBetweenV2(head *ListNode, left int, right int) *ListNode {
	dummyNode := &ListNode{Next: head}
	// 前驱
	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	// left 这里.....
	// todo 锻炼下loop反转链表
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	// 这个也没啥问题
	return dummyNode.Next
}
