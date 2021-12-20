package linked_list

import (
	"fmt"
	"testing"
)

// ["FrontMiddleBackQueue","pushMiddle","pushMiddle","pushMiddle","popMiddle","popMiddle","popMiddle"]
// [[],[1],[2],[3],[],[],[]]

func TestFrontMiddleBackQueueConstructor(t *testing.T) {
	ff := FrontMiddleBackQueueConstructor()
	ff.PushMiddle(1)
	ff.PushMiddle(2)
	ff.PushMiddle(3)
	// 1
	// 21
	// 231
	i := ff.PopMiddle()
	fmt.Print(i)
	i1 := ff.PopMiddle()
	fmt.Print(i1)
	i2 := ff.PopMiddle()
	fmt.Print(i2)

}
