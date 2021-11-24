package linked_list

// 剑指 Offer 06. 从尾到头打印链表
// https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
func reversePrint(head *ListNode) []int {
	var ans []int
	var down func(*ListNode)
	down = func(node *ListNode) {
		if nil == node {
			return
		}
		down(node.Next)
		ans = append(ans, node.Val)
	}
	down(head)
	return ans
}
