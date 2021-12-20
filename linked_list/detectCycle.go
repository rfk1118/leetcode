package linked_list

// https://leetcode-cn.com/problems/linked-list-cycle-ii/
// 142. 环形链表 II
// 走一圈碰到的第一个就是
func detectCycle(head *ListNode) *ListNode {
	if nil == head {
		return head
	}
	m := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := m[head]; ok {
			return head
		}
		m[head] = struct{}{}
		head = head.Next
	}
	return nil
}

// 快慢指针先走，如果走到相遇的位置，让快速指针在从头走，再次相遇就是第一个节点
// 2(a+b)= a+n(b+c)+b
// a = nb - b + nc
// a = (n-1)(b+c) + c
// 继续走a,这边走c
func detectCycleV2(head *ListNode) *ListNode {
	if nil == head {
		return head
	}
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
