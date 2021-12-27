package top

import (
	"fmt"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	s := generateParenthesis(4)
	fmt.Print(s)
}
