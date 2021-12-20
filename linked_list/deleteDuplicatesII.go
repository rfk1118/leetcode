package linked_list

import "math"

// 82. 删除排序链表中的重复元素 II
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/
func deleteDuplicatesII(head *ListNode) *ListNode {
	if nil == head {
		return head
	}
	c := head
	m := make(map[int]int, 0)
	for c != nil {
		m[c.Val]++
		c = c.Next
	}
	dump := &ListNode{Val: math.MaxInt64, Next: head}
	current := dump
	for current.Next != nil {
		if m[current.Next.Val] > 1 {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return dump.Next
}

// 多重循环
func deleteDuplicatesIII(head *ListNode) *ListNode {
	if nil == head {
		return head
	}
	dump := &ListNode{Val: math.MaxInt64, Next: head}
	current := dump
	for current.Next != nil && current.Next.Next != nil {
		if current.Next.Val == current.Next.Next.Val {
			x := current.Next.Val
			// 这里会有一个多重循环
			for current.Next != nil && current.Next.Val == x {
				current.Next = current.Next.Next
			}
		} else {
			current = current.Next
		}
	}
	return dump.Next
}
