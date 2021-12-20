package leetcode

import (
	"math"
)

// https://leetcode-cn.com/problems/diameter-of-binary-tree/solution/er-cha-shu-de-zhi-jing-by-leetcode-solution/
// 543. 二叉树的直径
func diameterOfBinaryTree(root *TreeNode) int {
	diameter := math.MinInt64
	var maxDepth func(node *TreeNode) int
	maxDepth = func(node *TreeNode) int {
		if nil == node {
			return 0
		}
		lDepth := maxDepth(node.Left)
		rDepth := maxDepth(node.Right)
		diameter = Max(diameter, lDepth+rDepth+1)
		return Max(lDepth, rDepth) + 1
	}
	maxDepth(root)
	return diameter - 1
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
