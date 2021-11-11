package leetcode

// https://leetcode-cn.com/problems/merge-two-binary-trees/solution/he-bing-er-cha-shu-by-leetcode-solution/
// 617. 合并二叉树 dfs
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val = root1.Val + root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}

// bfs的代码好多啊，哈哈哈哈，但是使用的是queue处理
func mergeTrees02(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	// result root
	res := &TreeNode{Val: root1.Val + root2.Val}

	// result queue
	resQ := []*TreeNode{res}
	// bfs queue
	quR1 := []*TreeNode{root1}
	quR2 := []*TreeNode{root2}
	for len(quR1) > 0 && len(quR2) > 0 {
		resPop := resQ[0]
		resQ = resQ[1:]

		quR1Pop := quR1[0]
		quR1 = quR1[1:]
		quR2Pop := quR2[0]
		quR2 = quR2[1:]
		qR1L, qR1R := quR1Pop.Left, quR1Pop.Right
		qR2L, qR2R := quR2Pop.Left, quR2Pop.Right
		if qR1L != nil || qR2L != nil {
			if qR1L != nil && qR2L != nil {
				resPop.Left = &TreeNode{Val: qR1L.Val + qR2L.Val}
			} else if qR1L != nil {
				resPop.Left = qR1L
			} else {
				resPop.Left = qR2L
			}
		}
		if qR1R != nil || qR2R != nil {
			if qR1R != nil && qR2R != nil {
				resPop.Right = &TreeNode{Val: qR1R.Val + qR2R.Val}
			} else if qR1R != nil {
				resPop.Right = qR1R
			} else {
				resPop.Right = qR2R
			}
		}

	}

	return res
}
