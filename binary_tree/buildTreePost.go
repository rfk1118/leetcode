package leetcode

// 处理树
// l m r
// l r m
// 106. 从中序与后序遍历序列构造二叉树
// https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
func buildTreePost(inorder []int, postorder []int) *TreeNode {
	if nil == inorder || len(inorder) == 0 || nil == postorder || len(postorder) == 0 {
		return nil
	}
	length := len(postorder) - 1
	rootValue := postorder[length]
	root := &TreeNode{Val: rootValue}
	inorderMidIndex := 0
	for index, value := range inorder {
		if rootValue == value {
			inorderMidIndex = index
			break
		}
	}
	leftCount := len(inorder[:inorderMidIndex])
	root.Left = buildTreePost(inorder[:leftCount], postorder[:leftCount])
	root.Right = buildTreePost(inorder[leftCount+1:], postorder[leftCount:len(postorder)-1])
	return root
}

// todo while循环进行处理
func buildTreePostWithHash(inorder []int, postorder []int) *TreeNode {
	// key is inorder value
	// value is index
	intMap := map[int]int{}
	for index, value := range inorder {
		intMap[value] = index
	}
	var build func(int, int) *TreeNode
	build = func(inorderLeft, inorderRight int) *TreeNode {
		if inorderLeft > inorderRight {
			return nil
		}
		rootVal := postorder[len(postorder)-1]
		postorder = postorder[:len(postorder)-1]
		root := &TreeNode{Val: rootVal}
		rootIndex := intMap[rootVal]
		// 必须先构建右边
		root.Right = build(rootIndex+1, inorderRight)
		root.Left = build(inorderLeft, rootIndex-1)
		return root
	}
	return build(0, len(inorder)-1)
}

// todo 还是不会写啊
func buildTreePostWithLoop(inorder []int, postorder []int) *TreeNode {
	if nil == inorder || len(inorder) == 0 || nil == postorder || len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	stack := []*TreeNode{root}
	inorderIndex := len(inorder) - 1
	for i := len(postorder) - 2; i >= 0; i-- {
		postorderVal := postorder[i]
		node := stack[len(stack)-1]
		if node.Val != inorder[inorderIndex] {
			node.Right = &TreeNode{Val: postorderVal}
			stack = append(stack, node.Right)
		} else {
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex--
			}
			node.Left = &TreeNode{Val: postorderVal}
			stack = append(stack, node.Left)
		}
	}

	return root
}
