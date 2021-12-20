package linked_list

import (
	"fmt"
	"testing"
)

func TestStackOfPlatesConstructor(t *testing.T) {
	sop := StackOfPlatesConstructor(1)
	sop.Push(1)
	sop.Push(2)
	i := sop.PopAt(1)
	fmt.Print(i)
	i2 := sop.Pop()
	i3 := sop.Pop()
	fmt.Print(i2)
	fmt.Print(i3)
}
