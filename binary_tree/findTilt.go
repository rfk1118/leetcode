package leetcode

import "renfakai.com/leetcode/util"

// 563. 二叉树的坡度
// https://leetcode-cn.com/problems/binary-tree-tilt/
func findTilt(root *TreeNode) int {
	tilt := 0
	// 其实计算每个子树的和，然后走到mid时候计算一下就好
	var sumSubTree func(node *TreeNode) int
	sumSubTree = func(node *TreeNode) int {
		if nil == node {
			return 0
		}
		leftSum := sumSubTree(node.Left)
		rightSum := sumSubTree(node.Right)
		tilt = tilt + util.Abs(leftSum, rightSum)
		return node.Val + leftSum + rightSum
	}
	sumSubTree(root)
	return tilt
}
