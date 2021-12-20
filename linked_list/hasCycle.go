package linked_list

// https://leetcode-cn.com/problems/linked-list-cycle/
// 141. 环形链表

func hasCycle(head *ListNode) bool {
	if nil == head || head.Next == nil {
		return false
	}
	quick := head.Next
	slow := head
	for quick != slow {
		if quick == nil || quick.Next == nil {
			return false
		}
		quick = quick.Next.Next
		slow = slow.Next
	}
	return true
}
