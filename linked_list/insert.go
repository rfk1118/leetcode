package linked_list

// 剑指 Offer II 029. 排序的循环链表
// https://leetcode-cn.com/problems/4ueAj6/
func insert(aNode *Node, x int) *Node {
	//特殊判断
	if aNode == nil {
		node := &Node{Val: x}
		node.Next = node
		return node
	}

	//找到头尾节点
	max, min := aNode, aNode
	tmp := aNode.Next
	for tmp != aNode {
		if max.Val < tmp.Val {
			max = tmp
		}
		if min.Val > tmp.Val {
			min = tmp
		}
		tmp = tmp.Next
	}
	//若极大值==极小值，说明所有节点值都相同，直接插在头结点后面；
	if max.Val == min.Val {
		max.Next = &Node{Val: x, Next: max.Next}
	} else if max.Val <= x || min.Val >= x {
		//去重
		for max.Next.Val == max.Val {
			max = max.Next
		}
		max.Next = &Node{Val: x, Next: max.Next}
	} else {
		for min.Next.Val < x {
			min = min.Next
		}
		min.Next = &Node{Val: x, Next: min.Next}
	}
	return aNode
}
