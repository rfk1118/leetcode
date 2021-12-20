package linked_list

// https://leetcode-cn.com/problems/reorder-list/
// 143. 重排链表
func reorderList(head *ListNode) {
	var findMid func(node *ListNode) *ListNode
	findMid = func(node *ListNode) *ListNode {
		if nil == head {
			return head
		}
		quick, slow := head, head
		for quick != nil && quick.Next != nil {
			quick = quick.Next.Next
			slow = slow.Next
		}
		return slow
	}
	var reverse func(node *ListNode) *ListNode
	reverse = func(node *ListNode) *ListNode {
		if nil == node || node.Next == nil {
			return node
		}
		root := reverse(node.Next)
		node.Next.Next = node
		node.Next = nil
		return root
	}
	// todo 重点复习
	var merge func(*ListNode, *ListNode)
	merge = func(l1, l2 *ListNode) {
		var l1Tmp, l2Tmp *ListNode
		for l1 != nil && l2 != nil {
			// 拿到两个后继续
			l1Tmp = l1.Next
			l2Tmp = l2.Next
			// l1后面设置l2
			l1.Next = l2
			// 拿到后继续
			l1 = l1Tmp
			l2.Next = l1
			l2 = l2Tmp
		}
	}

	mid := findMid(head)
	l1 := head
	l2 := reverse(mid.Next)
	mid.Next = nil
	merge(l1, l2)
}

func reorderListV2(head *ListNode) {
	var t []*ListNode
	for c := head; c != nil; c = c.Next {
		t = append(t, c)
	}

	i, j := 0, len(t)-1
	for i < j {
		t[i].Next = t[j]
		i++
		if i == j {
			break
		}
		t[j].Next = t[i]
		j--
	}
	// todo ???
	t[i].Next = nil
}
