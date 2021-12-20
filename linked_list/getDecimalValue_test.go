package linked_list

import (
	"fmt"
	"testing"
)

func Test_getDecimalValue(t *testing.T) {
	value := getDecimalValue(&ListNode{Val: 0})
	fmt.Println(value)
}
