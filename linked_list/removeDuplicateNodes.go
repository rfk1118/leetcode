package linked_list

// 面试题 02.01. 移除重复节点
// https://leetcode-cn.com/problems/remove-duplicate-node-lcci/
func removeDuplicateNodes(head *ListNode) *ListNode {
	if nil == head {
		return head
	}
	m := map[int]bool{head.Val: true}
	current := head
	for current.Next != nil {
		if m[current.Next.Val] {
			current.Next = current.Next.Next
		} else {
			m[current.Next.Val] = true
			current = current.Next
		}
	}
	return head
}

func removeDuplicateNodesV2(head *ListNode) *ListNode {
	ob := head
	for ob != nil {
		// 移除所有与当前节点相同的数据
		oc := ob
		for oc.Next != nil {
			if oc.Next.Val == ob.Val {
				oc = oc.Next.Next
			} else {
				oc = oc.Next
			}
		}
		ob = ob.Next
	}
	return head
}
