package linked_list

// 剑指 Offer 18. 删除链表的节点
// https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof/
// 题目保证链表中节点的值互不相同
func deleteNode18(head *ListNode, val int) *ListNode {
	if nil == head {
		return nil
	}
	if head.Val == val {
		return head.Next
	}
	current := head
	for current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}
