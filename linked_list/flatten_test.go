package linked_list

import (
	"testing"
)

// [1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]
func Test_flatten(t *testing.T) {
	a1 := &flattenNode{Val: 1}
	a2 := &flattenNode{Val: 2}
	a1.Next = a2
	a2.Prev = a1
	a3 := &flattenNode{Val: 3}
	a2.Next = a3
	a3.Prev = a2
	a4 := &flattenNode{Val: 4}
	a3.Next = a4
	a4.Prev = a3
	a5 := &flattenNode{Val: 5}
	a4.Next = a5
	a5.Prev = a5
	a6 := &flattenNode{Val: 6}
	a5.Next = a6
	a6.Prev = a5

	a7 := &flattenNode{Val: 7}
	a8 := &flattenNode{Val: 8}
	a3.Child = a7
	a7.Next = a8
	a8.Prev = a7

	a9 := &flattenNode{Val: 9}
	a8.Next = a9
	a9.Prev = a8
	a10 := &flattenNode{Val: 10}
	a9.Next = a10
	a10.Prev = a9
	a11 := &flattenNode{Val: 11}
	a8.Child = a11
	a12 := &flattenNode{Val: 12}
	a11.Next = a12
	a12.Prev = a11
	flatten(a1)
}
