package top

// 2. 两数相加
// https://leetcode-cn.com/problems/add-two-numbers/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	add := 0
	dump := &ListNode{}
	c := dump
	for x, y := l1, l2; x != nil || y != nil || add != 0; {
		if x == nil && y == nil {
			c.Next = &ListNode{Val: add}
			break
		} else if x != nil && y != nil {
			c.Next = &ListNode{Val: add + x.Val + y.Val}
			x = x.Next
			y = y.Next
		} else if x != nil {
			c.Next = &ListNode{Val: add + x.Val}
			x = x.Next
		} else if y != nil {
			c.Next = &ListNode{Val: add + y.Val}
			y = y.Next
		}

		add, c.Next.Val = c.Next.Val/10, c.Next.Val%10
		c = c.Next

	}
	return dump.Next
}

func addTwoNumbersV2(l1 *ListNode, l2 *ListNode) *ListNode {
	for x, y := l1, l2; x != nil && y != nil; {
		if x.Next != nil && y.Next == nil {
			y.Next = &ListNode{Val: 0}
		} else if x.Next == nil && y.Next != nil {
			x.Next = &ListNode{Val: 0}
		}
		x = x.Next
		y = y.Next
	}
	dump := &ListNode{}
	c := dump
	add := 0
	for x, y := l1, l2; add != 0 || x != nil; {
		if x == nil {
			c.Next = &ListNode{Val: add}
			break
		}
		c.Next = &ListNode{Val: x.Val + y.Val + add}
		c.Next.Val, add = c.Next.Val%10, c.Next.Val/10
		c = c.Next
		x = x.Next
		y = y.Next
	}
	return dump.Next
}
