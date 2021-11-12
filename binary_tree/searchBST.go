package leetcode

// https://leetcode-cn.com/problems/search-in-a-binary-search-tree/
// 700. 二叉搜索树中的搜索
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	}
	return searchBST(root.Right, val)
}
