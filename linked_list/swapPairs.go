package linked_list

// 你能在不修改链表节点值的情况下解决这个问题吗?（也就是说，仅修改节点本身。）
// 每2个反转1次
// https://leetcode-cn.com/problems/swap-nodes-in-pairs/
// 24. 两两交换链表中的节点
// 第一想法是每几个为一组反转数据
func swapPairs(head *ListNode) *ListNode {
	if nil == head || head.Next == nil {
		return head
	}
	// 新头
	newHead := head.Next
	// 第2个节点后面继续处理
	// 递归处理
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}

func swapPairsLoop(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil {
		// 拿到头节点
		node1 := temp.Next
		node2 := temp.Next.Next
		// 处理头和下面两个数据的关系
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		// 移动2个位置
		temp = node1
	}
	return dummyHead.Next
}

//  分组？我的第一想法，一会做处理
func swapPairsGroup(head *ListNode) *ListNode {
	var reverse func(*ListNode, *ListNode) *ListNode
	reverse = func(a *ListNode, b *ListNode) *ListNode {
		if a.Next == b {
			return a
		}
		node := reverse(a.Next, b)
		a.Next.Next = a
		a.Next = b
		return node
	}
	var reverseKGroup func(node *ListNode, k int) *ListNode
	reverseKGroup = func(node *ListNode, k int) *ListNode {
		if nil == node {
			return node
		}
		// 区间 [a, b) 包含 k 个待反转元素
		a, b := node, node
		for i := 0; i < k; i++ {
			if nil == b {
				return node
			}
			b = b.Next
		}
		// [a,b] -> [b,a]
		newHead := reverse(a, b)
		a.Next = reverseKGroup(b, k)
		return newHead
	}
	return reverseKGroup(head, 2)
}
