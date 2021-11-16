package leetcode

// todo error
// 处理树
// l m r
// l r m
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
	root.Right = buildTreePost(inorder[leftCount:], postorder[leftCount:])
	return root
}
