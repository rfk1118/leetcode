package leetcode

// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
// 235. 二叉搜索树的最近公共祖先
// 一定要利用二叉搜索树的性质
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if nil == root {
		return root
	}
	if root.Val > q.Val && root.Val > p.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if root.Val < q.Val && root.Val < p.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}
