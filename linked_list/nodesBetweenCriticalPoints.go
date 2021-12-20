package linked_list

import "math"

// 2058. 找出临界点之间的最小和最大距离
// https://leetcode-cn.com/problems/find-the-minimum-and-maximum-number-of-nodes-between-critical-points/
func nodesBetweenCriticalPoints(head *ListNode) []int {
	var tmp []int
	prev := head
	tm := 0
	for c := head.Next; c != nil; c = c.Next {
		tm++
		if c.Next != nil {
			if c.Val < prev.Val && c.Val < c.Next.Val {
				tmp = append(tmp, tm)
			}
			if c.Val > prev.Val && c.Val > c.Next.Val {
				tmp = append(tmp, tm)
			}
		}
		// 提交了n次都因为少这一行error
		prev = prev.Next
	}

	if len(tmp) < 2 {
		return []int{-1, -1}
	} else if len(tmp) == 2 {
		return []int{tmp[1] - tmp[0], tmp[1] - tmp[0]}
	}
	min := math.MaxInt64
	for i := 0; i < len(tmp)-1; i++ {
		if tmp[i+1]-tmp[i] < min {
			min = tmp[i+1] - tmp[i]
		}
	}
	return []int{min, tmp[len(tmp)-1] - tmp[0]}
}
