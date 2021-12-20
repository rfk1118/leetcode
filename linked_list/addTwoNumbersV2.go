package linked_list

// 445. 两数相加 II
// https://leetcode-cn.com/problems/add-two-numbers-ii/
// 最佳结果使用栈
func addTwoNumbersV2(l1 *ListNode, l2 *ListNode) *ListNode {
	var length func(node *ListNode) int
	length = func(node *ListNode) int {
		ans := 0
		for ; nil != node; node = node.Next {
			ans++
		}
		return ans
	}
	lex1 := length(l1)
	lex2 := length(l2)
	var merge func(x, y *ListNode) (*ListNode, int)
	merge = func(x, y *ListNode) (*ListNode, int) {
		if x == nil || y == nil {
			return nil, 0
		}
		listNode, i := merge(x.Next, y.Next)
		node := &ListNode{Val: x.Val + y.Val + i}
		add := node.Val / 10
		node.Val = node.Val % 10
		node.Next = listNode
		return node, add
	}
	var handler func(l1 *ListNode, l2 *ListNode, l1C, l2C int) (*ListNode, int)
	handler = func(l1 *ListNode, l2 *ListNode, l1C, l2C int) (*ListNode, int) {
		if l1C > l2C {
			l1C--
			listNode, i := handler(l1.Next, l2, l1C, l2C)
			node := &ListNode{Val: l1.Val + i}
			add := node.Val / 10
			node.Val = node.Val % 10
			node.Next = listNode
			return node, add
		} else if l2C > l1C {
			l2C--
			listNode, i := handler(l1, l2.Next, l1C, l2C)
			node := &ListNode{Val: l2.Val + i}
			add := node.Val / 10
			node.Val = node.Val % 10
			node.Next = listNode
			return node, add
		} else {
			listNode, i := merge(l1.Next, l2.Next)
			node := &ListNode{Val: l1.Val + l2.Val + i}
			add := node.Val / 10
			node.Val = node.Val % 10
			node.Next = listNode
			return node, add
		}
	}

	node, i := handler(l1, l2, lex1, lex2)
	if i > 0 {
		root := &ListNode{Val: i}
		root.Next = node
		return root
	}
	return node
}

func addTwoNumbersV3(l1 *ListNode, l2 *ListNode) *ListNode {
	var stackL1 []*ListNode
	var stackL2 []*ListNode
	for c := l1; c != nil; c = c.Next {
		stackL1 = append(stackL1, c)
	}
	for c := l2; c != nil; c = c.Next {
		stackL2 = append(stackL2, c)
	}

	carry := 0
	var lastNode *ListNode
	for len(stackL1) > 0 || len(stackL2) > 0 || carry != 0 {
		a := 0
		if len(stackL1) != 0 {
			a = stackL1[len(stackL1)-1].Val
			stackL1 = stackL1[:len(stackL1)-1]
		}
		b := 0
		if len(stackL2) != 0 {
			b = stackL2[len(stackL2)-1].Val
			stackL2 = stackL2[:len(stackL2)-1]
		}

		node := &ListNode{Val: a + b + carry}
		carry = node.Val / 10
		node.Val = node.Val % 10
		node.Next = lastNode
		lastNode = node
	}
	return lastNode
}
