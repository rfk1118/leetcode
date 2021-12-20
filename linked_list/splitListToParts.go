package linked_list

// 725. 分隔链表
// https://leetcode-cn.com/problems/split-linked-list-in-parts/
// If there are N nodes in the list, and k parts, then every part has N/k elements, except the first N%k parts have an extra one.
func splitListToParts(head *ListNode, k int) []*ListNode {
	// ans
	parts := make([]*ListNode, k)
	// length
	l := 0
	for c := head; c != nil; c = c.Next {
		l++
	}
	// 每一组数量，N%k parts have an extra one
	quotient, remainder := l/k, l%k
	for i, curr := 0, head; i < k && curr != nil; i++ {
		parts[i] = curr
		partSize := quotient
		// 提前走一步extra one
		if i < remainder {
			partSize++
		}
		for j := 1; j < partSize; j++ {
			curr = curr.Next
		}
		curr, curr.Next = curr.Next, nil
	}
	return parts
}
