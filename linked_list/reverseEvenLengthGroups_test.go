package linked_list

import (
	"fmt"
	"testing"
)

func Test_reverseEvenLengthGroups(t *testing.T) {
	ln := generateList([]int{5, 2, 6, 3, 9, 1, 7, 3, 8, 4})
	ln2 := reverseEvenLengthGroups(ln)
	fmt.Print(ln2)
}
