package top

import (
	"fmt"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	s := longestPalindrome("babad")
	fmt.Println(s)
}

func Test_longestPalindromev3(t *testing.T) {
	s := longestPalindromev3("babad")
	fmt.Println(s)
}
