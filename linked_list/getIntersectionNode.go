package linked_list

// 160. 相交链表
// https://leetcode-cn.com/problems/intersection-of-two-linked-lists/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var length func(*ListNode) int
	// 题目数据 保证 整个链式结构中不存在环。
	length = func(node *ListNode) int {
		if node == nil {
			return 0
		}
		res := 0
		for node != nil {
			res++
			node = node.Next
		}
		return res
	}
	lA := length(headA)
	lB := length(headB)
	var goNStep func(nodeA, nodeB *ListNode, step int) *ListNode
	goNStep = func(nodeA, nodeB *ListNode, step int) *ListNode {
		for ; step > 0; step-- {
			nodeA = nodeA.Next
		}
		for nodeA != nil && nodeB != nil {
			if nodeA == nodeB {
				return nodeB
			} else {
				nodeA = nodeA.Next
				nodeB = nodeB.Next
			}
		}
		return nil
	}
	if lA > lB {
		return goNStep(headA, headB, lA-lB)
	}
	return goNStep(headB, headA, lB-lA)
}

func getIntersectionNodeMap(headA, headB *ListNode) *ListNode {
	vis := map[*ListNode]bool{}
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		vis[tmp] = true
	}
	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if vis[tmp] {
			return tmp
		}
	}
	return nil
}

// todo 太秀了
func getIntersectionNodeWithOther(headA, headB *ListNode) *ListNode {
	if nil == headA || nil == headB {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}
