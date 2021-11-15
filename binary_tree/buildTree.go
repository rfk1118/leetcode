package leetcode

// 前序遍历 {root,[left],[right]}
// 中序遍历{[left],root,[right]}
// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
// 105. 从前序与中序遍历序列构造二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if nil == preorder || len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	mid := 0
	for i := range inorder {
		if preorder[0] == inorder[i] {
			mid = i
			break
		}
	}
	leftCount := len(inorder[:mid])
	root.Left = buildTree(preorder[1:leftCount+1], inorder[:mid])
	root.Right = buildTree(preorder[leftCount+1:], inorder[mid+1:])
	return root
}

// loop
////
//func buildTreeLoop(preorder []int, inorder []int) *TreeNode {
//
//}
