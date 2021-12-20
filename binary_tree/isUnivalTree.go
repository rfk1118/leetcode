package leetcode

// https://leetcode-cn.com/problems/univalued-binary-tree/
// 965. 单值二叉树
func isUnivalTree(root *TreeNode) bool {
	if nil == root {
		return true
	}
	var bfs func(node *TreeNode) bool
	bfs = func(node *TreeNode) bool {
		if nil == node {
			return true
		}
		return node.Val == root.Val && bfs(node.Left) && bfs(node.Right)
	}
	return bfs(root)
}
