package linked_list

// https://leetcode-cn.com/problems/odd-even-linked-list/solution/qi-ou-lian-biao-by-leetcode-solution/
// 328. 奇偶链表
func oddEvenList(head *ListNode) *ListNode {
	if nil == head {
		return head
	}
	odd := head
	evenHead := head.Next
	even := evenHead
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}
