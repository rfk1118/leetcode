package leetcode

// [99. 恢复二叉搜索树](https://leetcode-cn.com/problems/recover-binary-search-tree/)
func recoverTree(root *TreeNode) {
	var nums []int
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if nil == node {
			return
		}
		inorder(node.Left)
		nums = append(nums, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	x, y := findTwoSwapped(nums)
	recoverT(root, 2, x, y)
}

func recoverT(root *TreeNode, count, x, y int) {
	if nil == root {
		return
	}
	if root.Val == x || root.Val == y {
		if root.Val == x {
			root.Val = y
		} else {
			root.Val = x
		}
		count--
		if count == 0 {
			return
		}
	}
	recoverT(root.Left, count, x, y)
	recoverT(root.Right, count, x, y)
}

// 应该有两种结果
func findTwoSwapped(nums []int) (int, int) {
	index01, index02 := -1, -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			index02 = i + 1
			if index01 == -1 {
				// 处理第一个
				index01 = i
			} else {
				// 走到这里说明有两个元素不同
				break
			}
		}
	}
	x, y := nums[index01], nums[index02]
	return x, y
}

// 这种最坑
func recoverTreeWithStack(root *TreeNode) {
	var stack []*TreeNode
	var x, y, pred *TreeNode
	for len(stack) > 0 || root != nil {
		for nil != root {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pred != nil && root.Val < pred.Val {
			y = root
			if x == nil {
				x = pred
			} else {
				break
			}
		}
		pred = root
		root = root.Right
	}
	x.Val, y.Val = y.Val, x.Val
}

func recoverTreeMorris(root *TreeNode) {
	if nil == root {
		return
	}
	var x, y, pred, predecessor *TreeNode
	var current = root
	for current != nil {
		if current.Left == nil {
			// do Something
			if nil != pred && current.Val < pred.Val {
				y = current
				if nil == x {
					x = pred
				}
			}
			pred = current
			current = current.Right
		} else {
			// l - r - r
			predecessor = current.Left
			for predecessor.Right != nil && predecessor.Right != current {
				predecessor = predecessor.Right
			}
			if predecessor.Right == nil {
				predecessor.Right = current
				current = current.Left
			} else {
				//  do Something
				if pred != nil && current.Val < pred.Val {
					y = current
					if x == nil {
						x = pred
					}
				}
				predecessor.Right = nil
				pred = current
				current = current.Right
			}
		}
	}
	x.Val, y.Val = y.Val, x.Val
}

func recoverTreeDfs(root *TreeNode) {
	if nil == root {
		return
	}
	var x, y, pred *TreeNode
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pred != nil && node.Val < pred.Val {
			y = node
			if x == nil {
				x = pred
			}
		}
		pred = node
		dfs(node.Right)
	}
	dfs(root)
	x.Val, y.Val = y.Val, x.Val
}
