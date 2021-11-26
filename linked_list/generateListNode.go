package linked_list

func generateList(r []int) *ListNode {
	dump := &ListNode{}
	current := dump
	for _, i := range r {
		node := &ListNode{Val: i}
		current.Next = node
		current = current.Next
	}
	return dump.Next
}
