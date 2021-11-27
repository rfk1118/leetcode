package leetcode

// flatten
// https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/
// 114. 二叉树展开为链表
// 这个题不是增长树那个
func flatten(root *TreeNode) {
	var l []*TreeNode
	var preHandler func(*TreeNode)
	preHandler = func(node *TreeNode) {
		if nil == node {
			return
		}
		l = append(l, node)
		preHandler(node.Left)
		preHandler(node.Right)
	}
	preHandler(root)
	for i := 1; i < len(l); i++ {
		pre, cur := l[i-1], l[i]
		pre.Left, pre.Right = nil, cur
	}
}

// todo 还有两种方法还没看
