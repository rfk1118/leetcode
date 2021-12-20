package linked_list

// https://leetcode-cn.com/problems/sort-list/
// 148. 排序链表
// 归并排序吗
// 常数级空间复杂度下肯定不能用数组在分裂
// O(n log n) 常规算法时间复杂度都不行
func merge(sort *ListNode, sort2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, sort, sort2
	for temp1 != nil && temp2 != nil {
		if temp1.Val < temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	}
	if temp2 != nil {
		temp.Next = temp2
	}
	return dummyHead.Next
}
func sortList(head *ListNode) *ListNode {
	var sort func(head, tail *ListNode) *ListNode
	sort = func(head, tail *ListNode) *ListNode {
		if head == nil {
			return head
		}
		if head.Next == tail {
			head.Next = nil
			return head
		}
		fast, slow := head, head
		for fast != tail {
			slow = slow.Next
			fast = fast.Next
			if fast != tail {
				fast = fast.Next
			}
		}
		mid := slow
		return merge(sort(head, mid), sort(mid, tail))
	}
	return sort(head, nil)

}

func sortListV2(head *ListNode) *ListNode {
	if nil == head {
		return head
	}
	// get length
	length := 0
	for c := head; c != nil; c = c.Next {
		length++
	}
	// 哑节点
	dummyHead := &ListNode{Next: head}
	for subLength := 1; subLength < length; subLength <<= 1 {
		prev, cur := dummyHead, dummyHead.Next
		for cur != nil {
			head1 := cur
			for i := 1; i < subLength && cur.Next != nil; i++ {
				cur = cur.Next
			}
			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for i := 1; i < subLength && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}
			// 保存h1,h2后面的数据
			var next *ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}
			// 合并后的数据
			prev.Next = merge(head1, head2)
			// 向后推进
			for prev.Next != nil {
				prev = prev.Next
			}
			// 还原
			cur = next
		}
	}
	// 哑节点后面的数据
	return dummyHead.Next
}
