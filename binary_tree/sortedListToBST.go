package leetcode

// https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/submissions/
// 109. 有序链表转换二叉搜索树
func sortedListToBSTT(head *ListNode) *TreeNode {
	if nil == head {
		return nil
	}
	var findMid func(*ListNode, *ListNode) *ListNode
	findMid = func(left, right *ListNode) *ListNode {
		// q s find mid
		slow, quick := left, left
		for quick.Next != right && quick.Next.Next != right {
			slow = slow.Next
			quick = quick.Next.Next
		}
		return slow
	}
	var build func(*ListNode, *ListNode) *TreeNode
	build = func(head *ListNode, tail *ListNode) *TreeNode {
		if head == tail {
			return nil
		}
		mid := findMid(head, tail)
		if nil == mid {
			return nil
		}
		root := &TreeNode{Val: mid.Val}
		root.Left = build(head, mid)
		root.Right = build(mid.Next, tail)
		return root
	}
	return build(head, nil)
}

// 第二种方式
var globalHead *ListNode

// 超出时间限制
func sortedListToBST(head *ListNode) *TreeNode {
	globalHead = head
	length := getLength(head)
	var build func(int, int) *TreeNode
	build = func(s int, e int) *TreeNode {
		if s > e {
			return nil
		}
		mid := (s + e + 1) / 2
		root := &TreeNode{}
		root.Left = build(s, mid-1)
		root.Val = globalHead.Val
		globalHead = globalHead.Next
		root.Right = build(mid+1, e)
		return root
	}
	return build(0, length-1)
}

func getLength(head *ListNode) (res int) {
	for ; head != nil; head = head.Next {
		res++
	}
	return res
}
