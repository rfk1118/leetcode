package dp

import (
	"fmt"
	"testing"
)

func Test_cutRod(t *testing.T) {
	i := cutRod(p, 4)
	fmt.Print(i)
}

func Test_bottomUp(t *testing.T) {
	i := bottomUp(p, 4)
	fmt.Print(i)
}
