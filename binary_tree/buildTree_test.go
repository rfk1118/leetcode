package leetcode

import (
	"fmt"
	"testing"
)

// todo 核心复习
func Test_buildTreeLoop(t *testing.T) {
	preorder := []int{3, 9, 8, 5, 4, 10, 20, 15, 7}
	inorder := []int{4, 5, 8, 10, 9, 3, 15, 20, 7}
	buildTreeLoop(preorder, inorder)
}

func Test_buildTree(t *testing.T) {
	preorder := []int{3, 9, 8, 5, 4, 10, 20, 15, 7}
	inorder := []int{4, 5, 8, 10, 9, 3, 15, 20, 7}
	tree := buildTree(preorder, inorder)
	fmt.Println(tree)
}
