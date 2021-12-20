package leetcode

import "strconv"

// https://leetcode-cn.com/problems/binary-tree-paths/
// 257. 二叉树的所有路径
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	var res []string
	var binaryTreePathsWithPath func(node *TreeNode, path string)
	binaryTreePathsWithPath = func(node *TreeNode, path string) {
		if node.Left == nil && node.Right == nil {
			res = append(res, path)
			return
		}
		if node.Left != nil {
			newPath := path + "->" + strconv.Itoa(node.Left.Val)
			binaryTreePathsWithPath(node.Left, newPath)
		}
		if node.Right != nil {
			newPath := path + "->" + strconv.Itoa(node.Right.Val)
			binaryTreePathsWithPath(node.Right, newPath)
		}
	}
	binaryTreePathsWithPath(root, strconv.Itoa(root.Val))
	return res
}
