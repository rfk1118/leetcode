package leetcode

// https://leetcode-cn.com/problems/average-of-levels-in-binary-tree/
// 637. 二叉树的层平均值
func averageOfLevels(root *TreeNode) (r []float64) {
	if nil == root {
		return nil
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		length := len(q)
		sum := 0
		// 每次都处理一个层级
		for i := 0; i < length; i++ {
			node := q[i]
			sum = sum + node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		r = append(r, float64(sum)/float64(length))
		// 上一层级推出
		q = q[length:]
	}
	return r
}

// 这里使用深度的其实并不太好，dfs
var m map[int]*averageOfLevel

type averageOfLevel struct {
	value int
	count int
}

func averageOfLevels02(root *TreeNode) (r []float64) {
	m = make(map[int]*averageOfLevel, 0)
	averageOfLevelsHelper(root, 1)
	for i := 1; i <= len(m); i++ {
		a := m[i]
		r = append(
			r,
			float64(a.value)/float64(a.count),
		)
	}
	return r
}

func averageOfLevelsHelper(root *TreeNode, lev int) {
	if nil == root {
		return
	}
	v := m[lev]
	if v == nil {
		level := &averageOfLevel{
			value: root.Val,
			count: 1,
		}
		m[lev] = level
	} else {
		m[lev] = &averageOfLevel{
			value: v.value + root.Val,
			count: v.count + 1,
		}
	}
	averageOfLevelsHelper(root.Left, lev+1)
	averageOfLevelsHelper(root.Right, lev+1)
}
