package leetcode

import "renfakai.com/leetcode/util"

// https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
// 104. 二叉树的最大深度
// 自底而上处理问题
func maxDepth(root *TreeNode) int {
	if nil == root {
		return 0
	}
	var maxDep func(node *TreeNode) int
	maxDep = func(node *TreeNode) int {
		if nil == node {
			return 0
		}
		// leaf
		if nil == node.Left && nil == node.Right {
			return 1
		}
		leftD := maxDep(node.Left)
		rightD := maxDep(node.Right)
		return util.Max(leftD, rightD) + 1
	}
	return maxDep(root)
}
