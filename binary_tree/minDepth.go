package leetcode

import (
	"math"
	"renfakai.com/leetcode/util"
)

// https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/submissions/
// 111. 二叉树的最小深度
func minDepth(root *TreeNode) int {
	// 整个树是空的
	if nil == root {
		return 0
	}
	var minHDepth func(node *TreeNode) int
	minHDepth = func(node *TreeNode) int {
		// 必须是叶子节点,这个很关键
		if nil == node.Left && nil == node.Right {
			return 1
		}
		min := math.MaxInt64
		if nil != node.Left {
			min = util.Min(minHDepth(node.Left), min)
		}
		if nil != node.Right {
			min = util.Min(minHDepth(node.Right), min)
		}
		return min + 1
	}
	return minHDepth(root)
}
