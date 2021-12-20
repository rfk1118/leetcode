package linked_list

import (
	"math/rand"
	"time"
)

// Solution https://leetcode-cn.com/problems/linked-list-random-node/solution/xu-shui-chi-chou-yang-suan-fa-sha-zi-du-neng-kan-d/
// todo 我没看明白
// 382. 链表随机节点
// https://leetcode-cn.com/problems/linked-list-random-node/
type Solution struct {
	head *ListNode
	r    *rand.Rand
}

func SolutionConstructor(head *ListNode) Solution {
	return Solution{head: head, r: rand.New(rand.NewSource(time.Now().Unix()))}
}

func (this *Solution) GetRandom() int {
	p := this.head
	i := 1
	var res int
	for p != nil {
		if this.r.Intn(i) == 0 {
			res = p.Val
		}
		p = p.Next
		i++
	}
	return res
}
