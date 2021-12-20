package leetcode

import "renfakai.com/leetcode/util"

// https://leetcode-cn.com/problems/balanced-binary-tree/
// 110. 平衡二叉树
func isBalanced(root *TreeNode) bool {
	if nil == root {
		return true
	}
	var isBalancedHandler func(node *TreeNode) (bool, int)
	isBalancedHandler = func(node *TreeNode) (bool, int) {
		if node == nil {
			return true, 0
		}
		// 这个是自底而上进行处理，因为每次都把结果使用栈传递上来，所以性能会好一些
		leftIsB, leftC := isBalancedHandler(node.Left)
		rightISB, rightC := isBalancedHandler(node.Right)
		return leftIsB && rightISB && util.Abs(leftC, rightC) <= 1, util.Max(leftC, rightC) + 1
	}
	res, _ := isBalancedHandler(root)
	return res
}

// 下面的代码容易理解，但是对于相同的节点都进行处理过
//func isBalanced(root *TreeNode) bool {
//	if nil == root {
//		return true
//	}
//	return abs(height(root.Right), height(root.Left)) <= 1 &&
//		isBalanced(root.Left) &&
//		isBalanced(root.Right)
//}
//
//func height(root *TreeNode) int {
//	if nil == root {
//		return 0
//	}
//	return Max(height(root.Left), height(root.Right)) + 1
//}
//
//func abs(a, b int) int {
//	if a < b {
//		return b - a
//	} else {
//		return a - b
//	}
//}
//
//func Max(a, b int) int {
//	if a < b {
//		return b
//	}
//	return a
//}
