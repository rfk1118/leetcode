package linked_list

// https://leetcode-cn.com/problems/middle-of-the-linked-list/
// 876. 链表的中间结点
func middleNode(head *ListNode) *ListNode {
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

//
//func middleNode(head *ListNode) *ListNode {
//	if nil == head {
//		return head
//	}
//	quick, slow := head, head
// 对于中间节点为偶数的无法处理
//	for quick.Next != nil && quick.Next.Next != nil {
//		quick = quick.Next.Next
//		slow = slow.Next
//	}
//	return slow
//}
