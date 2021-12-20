package leetcode

import (
	"fmt"
	"testing"
)

func Test_pathSum(t *testing.T) {

	//  [5,4,8,11,null,13,4,7,2,null,null,5,1]
	//	22

	//node5 := &TreeNode{
	//	Val: 5,
	//}
	//node4 := &TreeNode{
	//	Val: 4,
	//}
	//node8 := &TreeNode{
	//	Val: 8,
	//}
	//node11 := &TreeNode{
	//	Val: 11,
	//}
	//node13 := &TreeNode{
	//	Val: 13,
	//}
	//node44 := &TreeNode{
	//	Val: 4,
	//}
	//node7 := &TreeNode{
	//	Val: 7,
	//}
	//node2 := &TreeNode{
	//	Val: 2,
	//}
	//node55 := &TreeNode{
	//	Val: 5,
	//}
	//node1 := &TreeNode{
	//	Val: 1,
	//}
	//node5.Left = node4
	//node5.Right = node8
	//node4.Left = node11
	//node11.Left = node7
	//node11.Right = node2
	//node8.Left = node13
	//node8.Right = node44
	//node44.Left = node55
	//node44.Right = node1
	node2 := &TreeNode{Val: -2}
	node3 := &TreeNode{Val: -3}
	node2.Right = node3
	sum := pathSum(node2, -5)
	fmt.Println(sum)
}
