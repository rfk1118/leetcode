package leetcode

// 938. 二叉搜索树的范围和
// https://leetcode-cn.com/problems/range-sum-of-bst/
// 二叉搜索树
func rangeSumBST(root *TreeNode, low int, high int) int {
	var handler func(node *TreeNode, low int, high int) int
	handler = func(node *TreeNode, low int, high int) int {
		if nil == node {
			return 0
		}
		if node.Val < low {
			return handler(node.Right, low, high)
		}
		if node.Val > high {
			return handler(node.Left, low, high)
		}
		return handler(node.Left, low, high) + handler(node.Right, low, high) + node.Val
	}
	return handler(root, low, high)
}

// 108 ms	9.4 MB
func rangeSumBST02(root *TreeNode, low int, high int) int {
	var res int
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		if node.Left != nil && node.Val > high {
			q = append(q, node.Left)
		}
		if node.Right != nil && node.Val < low {
			q = append(q, node.Right)
		}
		if node.Val >= low && node.Val <= high {
			res = res + node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		q = q[1:]
	}
	return res
}
