package linked_list

// 面试题 02.02. 返回倒数第 k 个节点
// https://leetcode-cn.com/problems/kth-node-from-end-of-list-lcci/
func kthToLast(head *ListNode, k int) int {
	if nil == head {
		return 0
	}
	var handler func(node *ListNode) (int, *ListNode)
	handler = func(node *ListNode) (int, *ListNode) {
		if node == nil {
			return 0, nil
		}
		last, listNode := handler(node.Next)
		if nil != listNode {
			return last, listNode
		}
		if last+1 == k {
			return last, node
		}
		return last + 1, nil
	}
	_, node := handler(head)
	return node.Val
}
