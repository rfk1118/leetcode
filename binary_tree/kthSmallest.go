package leetcode

// 二叉搜索树 所以使用中序遍历
// 230. 二叉搜索树中第K小的元素
// https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/
func kthSmallest(root *TreeNode, k int) int {
	var heap []int
	var kthSmallestHelper func(node *TreeNode)
	kthSmallestHelper = func(node *TreeNode) {
		if nil == node {
			return
		}
		kthSmallestHelper(node.Left)
		heap = append(heap, node.Val)
		kthSmallestHelper(node.Right)
	}
	kthSmallestHelper(root)
	if k > len(heap) {
		return 0
	}
	return heap[k-1]
}
