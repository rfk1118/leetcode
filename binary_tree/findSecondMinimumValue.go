package leetcode

// https://leetcode-cn.com/problems/second-minimum-node-in-a-binary-tree/solution/er-cha-shu-zhong-di-er-xiao-de-jie-dian-bhxiw/
// 671. 二叉树中第二小的节点
// 因为这里找值的时候会出现重复，所以不能使用常规方案
func findSecondMinimumValue(root *TreeNode) int {
	ans := -1
	var find func(node *TreeNode)
	find = func(node *TreeNode) {
		if nil == node {
			return
		}
		if ans != -1 && node.Val >= ans {
			return
		}
		if node.Val > root.Val {
			ans = node.Val
		}
		find(node.Left)
		find(node.Right)
	}
	find(root)
	return ans
}
