package linked_list

import (
	"fmt"
	"testing"
)

func Test_BrowserHistory(t *testing.T) {
	// 	["BrowserHistory","visit","visit","visit","back","back","forward","visit","forward","back","back"]
	// [["leetcode.com"],["google.com"],["facebook.com"],["youtube.com"],[1],[1],[1],["linkedin.com"],[2],[2],[7]]

	bh := BrowserHistoryConstructor("leetcode.com")
	bh.Visit("google.com")
	bh.Visit("facebook.com")
	bh.Visit("youtube.com")
	bh.Back(1)
	bh.Back(1)
	bh.Forward(1)
	bh.Visit("linkedin.com")
	bh.Forward(2)
	bh.Back(2)
	bh.Back(7)

	fmt.Print(bh)
}
