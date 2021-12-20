package leetcode

// 107. 二叉树的层序遍历 II
// https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/
func levelOrderBottom(root *TreeNode) [][]int {
	if nil == root {
		return nil
	}
	var res [][]int
	q := []*TreeNode{root}
	for len(q) > 0 {
		length := len(q)
		var tmp []int
		for i := 0; i < length; i++ {
			node := q[i]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, tmp)
		q = q[length:]
	}
	var resLen = len(res)
	for min := 0; min < resLen/2; min++ {
		res[min], res[resLen-1-min] = res[resLen-1-min], res[min]
	}
	return res
}
