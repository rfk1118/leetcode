package linked_list

// 面试题 04.03. 特定深度节点链表
// https://leetcode-cn.com/problems/list-of-depth-lcci/

func listOfDepth(tree *TreeNode) []*ListNode {
	var ans []*ListNode
	if nil == tree {
		return ans
	}
	q := []*TreeNode{tree}
	for len(q) > 0 {
		length := len(q)
		dump := &ListNode{}
		currency := dump
		for i := 0; i < length; i++ {
			tn := q[i]
			if nil != tn.Left {
				q = append(q, tn.Left)
			}
			if nil != tn.Right {
				q = append(q, tn.Right)
			}
			currency.Next = &ListNode{Val: tn.Val}
			currency = currency.Next
		}
		ans = append(ans, dump.Next)
		q = q[length:]
	}
	return ans
}
