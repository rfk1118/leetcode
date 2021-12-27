package top

// 21. 合并两个有序链表
// https://leetcode-cn.com/problems/merge-two-sorted-lists/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dump := &ListNode{}
	c, x, y := dump, list1, list2
	for x != nil && y != nil {
		if x.Val < y.Val {
			c.Next = x
			x = x.Next
		} else {
			c.Next = y
			y = y.Next
		}
		c = c.Next
	}
	if x == nil {
		c.Next = y
	}
	if y == nil {
		c.Next = x
	}
	return dump.Next
}
