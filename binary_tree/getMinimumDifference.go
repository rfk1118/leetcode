package leetcode

import (
	"math"
)

// https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/submissions/
// 530. 二叉搜索树的最小绝对差
//  这里使用了一个数组当作辅助空间，性能不好
// 二叉搜索树
func getMinimumDifference(root *TreeNode) int {
	var temp []int
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if nil == node {
			return
		}
		inorder(node.Left)
		temp = append(temp, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	length := len(temp)
	min := math.MaxInt64
	// 使用后继节点
	for i := range temp {
		if i < length && (i+1) < length {
			min = Min(min, temp[i+1]-temp[i])
		}
	}
	return min
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func getMinimumDifference02(root *TreeNode) int {
	res := math.MaxInt64
	pre := -1
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if nil == node {
			return
		}
		inorder(node.Left)
		if pre != -1 && node.Val-pre < res {
			res = node.Val - pre
		}
		pre = node.Val
		inorder(node.Right)
	}
	inorder(root)
	return res
}
