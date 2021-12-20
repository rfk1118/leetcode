package leetcode

// 103. 二叉树的锯齿形层序遍历
// https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/
// 果然是反转数组，如果是反转程序的话，代码会复杂很多
func zigzagLevelOrder(root *TreeNode) [][]int {
	if nil == root {
		return nil
	}
	var res = make([][]int, 0)
	var count = 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		length := len(q)
		var temp = make([]int, 0)
		for i := 0; i < length; i++ {
			node := q[i]
			temp = append(temp, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		//if count%2 != 0 {
		//	var temp02 []int
		//	for i := len(temp) - 1; i >= 0; i-- {
		//		temp02 = append(temp02, temp[i])
		//	}
		//	res = append(res, temp02)
		//} else {
		//	res = append(res, temp)
		//}
		if count%2 != 0 {
			for i, n := 0, len(temp); i < n/2; i++ {
				temp[i], temp[n-1-i] = temp[n-1-i], temp[i]
			}
		}
		res = append(res, temp)
		q = q[length:]
		count++
	}
	return res
}
