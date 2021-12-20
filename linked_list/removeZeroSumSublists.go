package linked_list

// 1171. 从链表中删去总和值为零的连续节点
// https://leetcode-cn.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list/
// 删去链表中由 总和 值为 0
// Convert the linked list into an array.
// While you can find a non-empty subarray with sum = 0, erase it.
// Convert the array into a linked list.
func removeZeroSumSublists(head *ListNode) *ListNode {
	if nil == head {
		return head
	}

	dumpNode := &ListNode{Val: 0, Next: head}
	m := make(map[int]*ListNode, 0)
	sum := 0
	// 首次遍历建立 节点处链表和<->节点 哈希表
	// 若同一和出现多次会覆盖，即记录该sum出现的最后一次节点
	for c := dumpNode; c != nil; c = c.Next {
		sum += c.Val
		m[sum] = c
	}

	sum = 0
	// 第二遍遍历 若当前节点处sum在下一处出现了则表明两结点之间所有节点和为0 直接删除区间所有节点
	for d := dumpNode; d != nil; d = d.Next {
		sum += d.Val
		d.Next = m[sum].Next
	}
	return dumpNode.Next
}
