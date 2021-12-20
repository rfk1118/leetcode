package leetcode

// https://leetcode-cn.com/problems/cousins-in-binary-tree/
// 993. 二叉树的堂兄弟节点
func isCousins(root *TreeNode, x int, y int) bool {
	var xParent, yParent *TreeNode
	var xDepth, yDepth int
	var xFound, yFound bool
	update := func(node, parent *TreeNode, level int) {
		if node == nil {
			return
		}
		if node.Val == x {
			xParent = parent
			xDepth = level
			xFound = true
		}
		if node.Val == y {
			yParent = parent
			yDepth = level
			yFound = true
		}
	}

	type Pair struct {
		node  *TreeNode
		depth int
	}

	q := []*Pair{{node: root, depth: 1}}
	for len(q) > 0 && !(xFound && yFound) {
		node := q[0]
		treeNode := node.node
		cDepth := node.depth
		if treeNode.Left != nil {
			update(treeNode.Left, treeNode, cDepth+1)
			q = append(q, &Pair{node: treeNode.Left, depth: cDepth + 1})
		}
		if treeNode.Right != nil {
			update(treeNode.Right, treeNode, cDepth+1)
			q = append(q, &Pair{node: treeNode.Right, depth: cDepth + 1})
		}
		q = q[1:]
	}
	return xDepth == yDepth && xParent != yParent
}
