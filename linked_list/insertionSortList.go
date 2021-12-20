package linked_list

// https://leetcode-cn.com/problems/insertion-sort-list/
// 147. 对链表进行插入排序
// 使用递归栈进行处理
// todo
func insertionSortList(head *ListNode) *ListNode {
	if nil == head {
		return head
	}
	var dummyHead = &ListNode{Next: head}
	// 最后一个排序的
	lastSorted, curr := head, head.Next
	for curr != nil {
		// 如果当前元素比已经排序的元素要大，往下走
		if lastSorted.Val <= curr.Val {
			lastSorted = lastSorted.Next
		} else {
			// 从头开始走，找到前驱
			prev := dummyHead
			for prev.Next.Val <= curr.Val {
				prev = prev.Next
			}
			// 未处理的挂到有序数组后面
			lastSorted.Next = curr.Next
			// 排序好的next,放到当前的后面
			curr.Next = prev.Next
			// 将当前放到prev后面
			prev.Next = curr
		}
		// 获取到未排序数据
		curr = lastSorted.Next
	}
	return dummyHead.Next
}
