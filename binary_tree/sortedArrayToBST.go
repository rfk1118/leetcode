package leetcode

// https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/
// 108. 将有序数组转换为二叉搜索树
// 处理问题使用的是自顶而下切分数组，可以查看归并排序
func sortedArrayToBST(nums []int) *TreeNode {
	if nil == nums || len(nums) == 0 {
		return nil
	}
	var handler func(nums []int, low, high int) *TreeNode
	handler = func(nums []int, low, high int) *TreeNode {
		if low > high {
			return nil
		}
		mid := (high-low)/2 + low
		res := &TreeNode{Val: nums[mid]}
		res.Left = handler(nums, low, mid-1)
		res.Right = handler(nums, mid+1, high)
		return res
	}
	return handler(nums, 0, len(nums)-1)
}
