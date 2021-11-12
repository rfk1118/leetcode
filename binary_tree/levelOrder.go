package leetcode

// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/
// 剑指 Offer 32 - II. 从上到下打印二叉树 II
// 每层求平均数的变种
func levelOrder(root *TreeNode) [][]int {
	if nil == root {
		return nil
	}
	var res = make([][]int, 0)
	q := []*TreeNode{root}
	for len(q) != 0 {
		length := len(q)
		var t []int
		for i := 0; i < length; i++ {
			node := q[i]
			t = append(t, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, t)
		q = q[length:]
	}
	return res
}
