package linked_list

// https://leetcode-cn.com/problems/convert-binary-number-in-a-linked-list-to-integer/
// 1290. 二进制链表转整数
func getDecimalValue(head *ListNode) int {
	ans := 0
	var handler func(*ListNode) int
	handler = func(node *ListNode) int {
		if node == nil {
			return ans
		}
		ans = ans<<1 + node.Val
		return handler(node.Next)
	}
	return handler(head)
}

// 搞笑这个方式居然通不过
//var ans int
//
//func getDecimalValue(head *ListNode) int {
//	if nil == head {
//		return ans
//	}
//	ans = ans<<1 + head.Val
//	return getDecimalValue(head.Next)
//}
