package linked_list

// 2074. 反转偶数长度组的节点
// https://leetcode-cn.com/problems/reverse-nodes-in-even-length-groups/
// 156 / 165 个通过测试用例 然后超出了时间限制
func reverseEvenLengthGroupsV2(head *ListNode) *ListNode {
	if nil == head {
		return head
	}
	var handler func(node *ListNode, count int) []*ListNode
	handler = func(node *ListNode, count int) []*ListNode {
		if nil == node {
			return nil
		}
		var tmp []*ListNode
		c := node
		for i := 0; i < count && c != nil; i++ {
			tmp = append(tmp, c)
			c = c.Next
		}
		if len(tmp)%2 == 0 {
			i, j := 0, len(tmp)-1
			for i <= j {
				tmp[i], tmp[j] = tmp[j], tmp[i]
				i++
				j--
			}
		}
		if c == nil {
			return tmp
		}
		ln := handler(c, count+1)
		tmp = append(tmp, ln...)
		return tmp
	}
	ln := handler(head, 1)
	for i := 0; i < len(ln)-1; i++ {
		ln[i].Next = ln[i+1]
	}
	ln[len(ln)-1].Next = nil
	return ln[0]
}

func reverseEvenLengthGroups(head *ListNode) *ListNode {
	var resverInner func(*ListNode, *ListNode) *ListNode
	resverInner = func(ln1, ln2 *ListNode) *ListNode {
		if ln1.Next == ln2 {
			return ln1
		}
		ln := resverInner(ln1.Next, ln2)
		ln1.Next.Next = ln1
		ln1.Next = ln2
		return ln
	}

	var handler func(node *ListNode, count int) *ListNode
	handler = func(node *ListNode, count int) *ListNode {
		if nil == node {
			return node
		}
		a, b, bPre, newNode := node, node, node, node
		tempCount := 0
		// 这里少走一步
		for i := 0; i < count && b != nil; i++ {
			tempCount++
			b = b.Next
			if i != count-1 {
				bPre = bPre.Next
			}
		}
		if tempCount%2 == 0 {
			newNode = resverInner(a, b)
			if b != nil {
				a.Next = handler(b, count+1)
			}
		} else {
			if b != nil {
				bPre.Next = handler(b, count+1)
			}
		}
		return newNode
	}
	return handler(head, 1)
}
