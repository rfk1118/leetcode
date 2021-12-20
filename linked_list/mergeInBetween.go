package linked_list

// 1669. 合并两个链表
// https://leetcode-cn.com/problems/merge-in-between-linked-lists/

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	if nil == list1 {
		return list1
	}
	var x, y *ListNode
	ans := 0
	for c := list1; c != nil; c = c.Next {
		if nil != x && y != nil {
			break
		}
		if ans == a-1 {
			x = c
		}
		if ans == b {
			y = c.Next
		}
		ans++
	}
	if nil == list2 {
		x.Next = y
		return list1
	}

	x.Next = list2
	c2 := list2
	for c2.Next != nil {
		c2 = c2.Next
	}
	c2.Next = y
	return list1
}
