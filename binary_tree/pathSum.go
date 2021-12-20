package leetcode

// https://leetcode-cn.com/problems/path-sum-ii/
// 113. 路径总和 II
func pathSum(root *TreeNode, targetSum int) [][]int {
	//if nil == root || root.Val > targetSum {
	//	return nil
	//}
	if nil == root {
		return nil
	}
	var res [][]int
	var bfs func(node *TreeNode, currentVal int, pre []int)
	bfs = func(node *TreeNode, currentVal int, pre []int) {
		//这里可能会有负数
		//if nil == node || currentVal > targetSum {
		//	return
		//}
		if nil == node {
			return
		}
		currentVal = currentVal + node.Val
		insert := append(pre, node.Val)
		if node.Left == nil && node.Right == nil {
			if targetSum == currentVal {
				// 被指针玩死了
				res = append(res, append([]int(nil), insert...))
			}
			return
		}
		if node.Left != nil {
			bfs(node.Left, currentVal, insert)
		}
		if node.Right != nil {
			bfs(node.Right, currentVal, insert)
		}
	}
	bfs(root, 0, []int{})
	return res
}
