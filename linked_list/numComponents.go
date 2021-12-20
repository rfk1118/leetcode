package linked_list

// 817. 链表组件
// https://leetcode-cn.com/problems/linked-list-components/
// 我居然想用union
func numComponents(head *ListNode, nums []int) int {
	var m = make(map[int]struct{}, 0)
	for _, v := range nums {
		m[v] = struct{}{}
	}
	ans := 0
	for c := head; c != nil; c = c.Next {
		if _, ok := m[c.Val]; ok {
			next := c.Next
			if nil == next {
				ans++
			} else {
				if _, ok := m[next.Val]; !ok {
					ans++
				}
			}
		}
	}
	return ans
}
