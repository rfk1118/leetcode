package leetcode

// https://leetcode-cn.com/problems/sum-of-left-leaves/submissions/
// 404. 左叶子之和
// 注意左子树的值不要算重复
func sumOfLeftLeaves(root *TreeNode) int {
	if nil == root {
		return 0
	}
	var dfs func(root *TreeNode) (res int)
	dfs = func(node *TreeNode) (res int) {
		q := []*TreeNode{node}
		// 非空
		for len(q) > 0 {
			pop := q[0]
			q = q[1:]
			if pop.Left != nil {
				// 叶子节点加入结果
				if isLeaf(pop.Left) {
					res = res + pop.Left.Val
				} else {
					q = append(q, pop.Left)
				}
			}
			if pop.Right != nil && !isLeaf(pop.Right) {
				res = res + dfs(pop.Right)
			}
		}
		return
	}
	return dfs(root)
}

func isLeaf(root *TreeNode) bool {
	return root.Left == nil && root.Right == nil
}
