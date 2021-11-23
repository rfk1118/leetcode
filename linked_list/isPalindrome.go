package linked_list

// https://leetcode-cn.com/problems/palindrome-linked-list/
// 234. 回文链表
// 变数组？
// 我第一反应是产生一个新的反转的链表
func isPalindrome(head *ListNode) bool {
	var temp []*ListNode
	for head != nil {
		temp = append(temp, head)
		head = head.Next
	}
	length := len(temp) - 1
	mid := length / 2
	for i := 0; i <= mid; i++ {
		if temp[i].Val != temp[length-i].Val {
			return false
		}
	}
	return true
}

func isPalindromeWithRecursively(head *ListNode) bool {
	front := head
	var recursively func(node *ListNode) bool
	recursively = func(node *ListNode) bool {
		if node != nil {
			// 一直向下走
			if !recursively(node.Next) {
				return false
			}
			// 对比
			if node.Next != front.Next {
				return false
			}
			// 设置头
			front = front.Next
		}
		return true
	}
	return recursively(head)
}

//整个流程可以分为以下五个步骤：
//找到前半部分链表的尾节点。
//反转后半部分链表。
//判断是否回文。
//恢复链表。
//返回结果。
func isPalindromeWithRecursivelyRe(head *ListNode) bool {
	var reverse func(node *ListNode) *ListNode
	reverse = func(node *ListNode) *ListNode {
		if nil == node || node.Next == nil {
			return node
		}
		root := reverse(node.Next)
		node.Next.Next = node
		node.Next = nil
		return root
	}
	var endOfFirstHalf func(node *ListNode) *ListNode
	endOfFirstHalf = func(node *ListNode) *ListNode {
		quick, slow := node, node
		for quick.Next != nil && quick.Next.Next != nil {
			quick = quick.Next.Next
			slow = slow.Next
		}
		return slow
	}
	if head == nil {
		return true
	}
	firstHalfEnd := endOfFirstHalf(head)
	// 反转
	secondHalfStart := reverse(firstHalfEnd.Next)
	p1 := head
	p2 := secondHalfStart
	result := true
	for result && p2 != nil {
		if p1.Val != p2.Val {
			result = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	firstHalfEnd.Next = reverseList(secondHalfStart)
	return result
}
