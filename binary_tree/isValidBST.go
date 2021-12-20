package leetcode

import "math"

// https://leetcode-cn.com/problems/validate-binary-search-tree/
// 98. 验证二叉搜索树
// 依然没写出来最好性能的版本
func isValidBST(root *TreeNode) bool {
	if nil == root {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	var findMin func(node *TreeNode) *TreeNode
	findMin = func(node *TreeNode) *TreeNode {
		if nil == node {
			return nil
		}
		if node.Left != nil {
			return findMin(node.Left)
		}
		return node
	}
	var findMax func(node *TreeNode) *TreeNode
	findMax = func(node *TreeNode) *TreeNode {
		if nil == node {
			return nil
		}
		if node.Right != nil {
			return findMax(node.Right)
		}
		return node
	}
	current := true
	min := findMin(root.Right)
	if min != nil {
		current = current && root.Val < min.Val
	}
	max := findMax(root.Left)
	if max != nil {
		current = current && root.Val > max.Val
	}
	return isValidBST(root.Left) && isValidBST(root.Right) && current
}

// 高性能版本
func isValidBSTH(root *TreeNode) bool {
	var validate func(node *TreeNode, min, max int) bool
	validate = func(node *TreeNode, min, max int) bool {
		if nil == node {
			return true
		}
		if node.Val <= min || node.Val >= max {
			return false
		}
		return validate(node.Left, min, node.Val) && validate(node.Right, node.Val, max)
	}
	return validate(root, math.MinInt64, math.MaxInt64)
}
