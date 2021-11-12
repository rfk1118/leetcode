package leetcode

// https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst/solution/liang-shu-zhi-he-iv-by-leetcode/
// 653. 两数之和 IV - 输入 BST
// dfs,bfs都用了hashmap,hashset也是hashmap.所以没什么大的区别
func findTarget(root *TreeNode, k int) bool {
	if nil == root {
		return false
	}
	var temp []int
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if nil == node {
			return
		}
		inorder(node.Left)
		temp = append(temp, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	// 双指针问题
	min, max := 0, len(temp)-1
	for min < max {
		if temp[min]+temp[max] == k {
			return true
		} else if temp[min]+temp[max] < k {
			min++
		} else {
			max--
		}
	}
	return false
}
