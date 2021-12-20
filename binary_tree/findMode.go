package leetcode

import "math"

// https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/solution/er-cha-sou-suo-shu-zhong-de-zhong-shu-by-leetcode-/
// 501. 二叉搜索树中的众数
func findMode(root *TreeNode) []int {
	if nil == root {
		return nil
	}
	// m store
	m := make(map[int]int, 0)
	var findModeHelper func(node *TreeNode)
	findModeHelper = func(node *TreeNode) {
		if node == nil {
			return
		}
		findModeHelper(node.Left)
		m[node.Val]++
		findModeHelper(node.Right)
	}
	findModeHelper(root)

	max := math.MinInt64
	for _, i := range m {
		if i > max {
			max = i
		}
	}

	var r []int
	for index, i := range m {
		if i == max {
			r = append(r, index)
		}
	}
	return r
}

func findMode02(root *TreeNode) []int {
	var base, count, maxCount int
	var answer []int
	update := func(x int) {
		// 当前值还没变化
		if x == base {
			count++
		} else {
			// 重置base
			base, count = x, 1
		}
		if count == maxCount {
			answer = append(answer, base)
		} else if count > maxCount {
			maxCount = count
			answer = []int{base}
		}
	}
	// 保证
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		update(node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return answer
}

// 方法二：Morris 中序遍历 hard?
// ????
func findMode03(root *TreeNode) (answer []int) {
	//  这里使用了正常流程
	var base, count, maxCount int
	update := func(x int) {
		if x == base {
			count++
		} else {
			base, count = x, 1
		}
		if count == maxCount {
			answer = append(answer, base)
		} else if count > maxCount {
			maxCount = count
			answer = []int{base}
		}
	}
	// init
	cur := root
	for cur != nil {
		// 如果当前节点没有左子树，则遍历这个点，然后跳转到当前节点的右子树。
		// 这里还有一种可能，那就是回溯到之前的节点
		if cur.Left == nil {
			update(cur.Val)
			cur = cur.Right
			continue
		}
		// 找到前驱，这里会有两种可能
		// 1.第一次找前驱
		// 2. 第二次找前驱
		pre := cur.Left
		for pre.Right != nil && pre.Right != cur {
			pre = pre.Right
		}
		// 第一次找前驱
		if pre.Right == nil {
			// 该前驱设置后继
			pre.Right = cur
			// 向下走
			cur = cur.Left
		} else {
			// 第二次找前驱，重置之前的连接
			pre.Right = nil
			update(cur.Val)
			cur = cur.Right
		}
	}
	return
}
