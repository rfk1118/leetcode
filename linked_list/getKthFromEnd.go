package linked_list

// 剑指 Offer 22. 链表中倒数第k个节点
// https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if nil == head {
		return head
	}
	var handler func(*ListNode) (int, *ListNode)
	handler = func(node *ListNode) (int, *ListNode) {
		if nil == node {
			return 0, nil
		}
		i, listNode := handler(node.Next)
		if nil != listNode {
			return i, listNode
		}
		if i+1 == k {
			return i + 1, node
		}
		return i + 1, nil
	}
	_, node := handler(head)
	return node
}
