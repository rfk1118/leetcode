package leetcode

// https://leetcode-cn.com/problems/sum-of-root-to-leaf-binary-numbers/submissions/
// 1022. 从根到叶的二进制数之和
func sumRootToLeaf(root *TreeNode) int {
	sum := 0
	var dfsSumRootToLeaf func(node *TreeNode, c int)
	dfsSumRootToLeaf = func(node *TreeNode, c int) {
		if nil == node {
			return
		}
		c = c<<1 + node.Val
		if node.Left == nil && node.Right == nil {
			sum += c
		}
		dfsSumRootToLeaf(node.Left, c)
		dfsSumRootToLeaf(node.Right, c)
	}
	dfsSumRootToLeaf(root, 0)
	return sum
}

func sumRootToLeaf02(root *TreeNode) int {
	type Pair struct {
		node *TreeNode
		c    int
	}
	var ans int
	q := []*Pair{{node: root, c: 0}}
	for len(q) > 0 {
		pair := q[0]
		node := pair.node
		c := pair.c
		c = c<<1 + node.Val
		if node.Left == nil && node.Right == nil {
			ans += c
		}
		if node.Left != nil {
			q = append(q, &Pair{node: node.Left, c: c})
		}
		if node.Right != nil {
			q = append(q, &Pair{node: node.Right, c: c})
		}
		q = q[1:]
	}
	return ans
}
