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
