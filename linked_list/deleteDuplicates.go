package linked_list

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
// 83. 删除排序链表中的重复元素
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	current := head
	for current.Next != nil {
		// 一直相等一直推进
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			// 不相等推其他数据
			current = current.Next
		}
	}
	return head
}
