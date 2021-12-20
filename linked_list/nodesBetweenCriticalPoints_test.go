package linked_list

import (
	"fmt"
	"testing"
)

func Test_nodesBetweenCriticalPoints(t *testing.T) {
	ln := generateList([]int{4, 2, 4, 1})
	i := nodesBetweenCriticalPoints(ln)
	fmt.Print(i)
}
