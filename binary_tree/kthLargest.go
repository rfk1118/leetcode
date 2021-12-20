package leetcode

// https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
// 剑指 Offer 54. 二叉搜索树的第k大节点
func kthLargest(root *TreeNode, k int) int {
	t := make([]int, 0)
	var orderKthLargest func(node *TreeNode)
	orderKthLargest = func(node *TreeNode) {
		if nil == node {
			return
		}
		orderKthLargest(node.Right)
		t = append(t, node.Val)
		orderKthLargest(node.Left)
	}
	orderKthLargest(root)
	return t[k-1]
}
