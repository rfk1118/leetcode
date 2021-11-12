package leetcode

import (
	"math"
)

// 783. 二叉搜索树节点最小距离
// https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/submissions/
// 差值是一个正数，其数值等于两值之差的绝对值。
// 差值必然为当前节点与前驱节点的值
func minDiffInBST(root *TreeNode) int {
	res := math.MaxInt64
	var pre *TreeNode = nil
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if nil == node {
			return
		}
		dfs(node.Left)
		if pre != nil {
			res = Min(res, Abs(node.Val, pre.Val))
		}
		pre = node
		dfs(node.Right)
	}
	dfs(root)
	return res
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
