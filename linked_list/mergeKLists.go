package linked_list

// https://leetcode-cn.com/problems/merge-k-sorted-lists/
// 23. 合并K个升序链表
//  这不是归并排序吗
// 一次过了
func mergeKLists(lists []*ListNode) *ListNode {
	if nil == lists || len(lists) == 0 {
		return nil
	}
	var merge func(*ListNode, *ListNode) *ListNode
	merge = func(node *ListNode, node2 *ListNode) *ListNode {
		if node == nil && node2 == nil {
			return nil
		}
		listNode := &ListNode{}
		current := listNode
		for node != nil || node2 != nil {
			if node == nil {
				current.Next = node2
				break
			} else if node2 == nil {
				current.Next = node
				break
			} else {
				if node.Val > node2.Val {
					current.Next = node2
					node2 = node2.Next
				} else {
					current.Next = node
					node = node.Next
				}
				current = current.Next
			}
		}
		return listNode.Next
	}
	var spilt func([]*ListNode) *ListNode
	spilt = func(nodes []*ListNode) *ListNode {
		if len(nodes) == 1 {
			return nodes[0]
		}
		i := len(nodes)
		node1 := spilt(nodes[0 : i/2])
		node2 := spilt(nodes[i/2:])
		return merge(node1, node2)
	}
	return spilt(lists)
}
