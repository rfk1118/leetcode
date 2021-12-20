package leetcode

import "strconv"

// https://leetcode-cn.com/problems/construct-string-from-binary-tree/
// 606. 根据二叉树创建字符串
func tree2str(root *TreeNode) string {
	if nil == root {
		return ""
	}
	// 如果当前节点没有孩子，那我们不需要在节点后面加上任何括号；
	if root.Left == nil && root.Right == nil {
		return strconv.Itoa(root.Val) + ""
	}
	// 如果当前节点只有左孩子，那我们在递归时，只需要在左孩子的结果外加上一层括号，而不需要给右孩子加上任何括号；
	if root.Right == nil {
		return strconv.Itoa(root.Val) + "(" + tree2str(root.Left) + ")"
	}
	// 如果当前节点只有右孩子，那我们在递归时，需要先加上一层空的括号 () 表示左孩子为空，
	// 再对右孩子进行递归，并在结果外加上一层括号。
	return strconv.Itoa(root.Val) + "(" + tree2str(root.Left) + ")" + "(" + tree2str(root.Right) + ")"
}
