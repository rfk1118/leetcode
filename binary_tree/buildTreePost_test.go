package leetcode

import (
	"fmt"
	"testing"
)

func Test_buildTreePost(t *testing.T) {
	//  [9,3,15,20,7]
	//	[9,15,7,20,3]
	post := buildTreePost([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	fmt.Println(post)
}

func Test_buildTreePostWithHash(t *testing.T) {
	post := buildTreePostWithHash([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	fmt.Println(post)
}
