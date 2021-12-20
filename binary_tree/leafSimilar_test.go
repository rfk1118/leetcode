package leetcode

import "testing"

func TestLeafSimilar(t *testing.T) {
	node1 := &TreeNode{
		Val: 1,
	}
	node2 := &TreeNode{
		Val: 2,
	}
	leafSimilar(node2, node1)
}
