package linked_list

// https://leetcode-cn.com/problems/next-greater-node-in-linked-list/
// 1019. 链表中的下一个更大节点

func nextLargerNodes(head *ListNode) []int {
	if nil == head {
		return nil
	}
	var ans = make([]int, 0)
	for c := head; c != nil; c = c.Next {
		d := c.Next
		flag := false
		for d != nil {
			if d.Val > c.Val {
				ans = append(ans, d.Val)
				flag = true
				break
			}
			d = d.Next
		}
		if !flag {
			ans = append(ans, 0)
		}
	}
	return ans
}

// error
// [1,7,5,1,9,2,5,1]
// [9,9,9,9,0,5,0,0]
// [7,9,9,9,0,5,0,0]
func nextLargerNodesV2(head *ListNode) []int {
	if nil == head {
		return nil
	}
	var ans []int = make([]int, 0)
	var handler func(node *ListNode) int
	handler = func(node *ListNode) int {
		if nil == node {
			return 0
		}
		i := handler(node.Next)
		if i > node.Val {
			ans = append(ans, i)
		} else {
			ans = append(ans, 0)
			i = node.Val
		}
		return i
	}
	handler(head)

	for i, j := 0, len(ans)-1; i <= j; {
		ans[i], ans[j] = ans[j], ans[i]
		i++
		j--
	}
	return ans
}
