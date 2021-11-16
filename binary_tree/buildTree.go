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

//loop
// 备注
func buildTreeLoop(preorder []int, inorder []int) *TreeNode {
	if nil == preorder {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	stack := []*TreeNode{root}
	var inorderIndex int
	// 必须剔除
	for index := 1; index < len(preorder); index++ {
		value := preorder[index]
		node := stack[len(stack)-1]
		if node.Val != inorder[inorderIndex] {
			node.Left = &TreeNode{value, nil, nil}
			stack = append(stack, node.Left)
		} else {
			for len(stack) != 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex++
			}
			node.Right = &TreeNode{value, nil, nil}
			stack = append(stack, node.Right)
		}
	}
	return root
}
