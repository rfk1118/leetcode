package linked_list

// 138. 复制带随机指针的链表
// https://leetcode-cn.com/problems/copy-list-with-random-pointer/

var cache map[*Node]*Node

func copyRandomList(head *Node) *Node {
	cache = map[*Node]*Node{}
	var deepCopy func(*Node) *Node
	deepCopy = func(node *Node) *Node {
		if nil == node {
			return node
		}
		if n, has := cache[node]; has {
			return n
		}
		newNode := &Node{Val: node.Val}
		cache[node] = newNode
		newNode.Next = deepCopy(node.Next)
		newNode.Random = deepCopy(node.Random)
		return newNode
	}
	return deepCopy(head)
}

func copyRandomListV2(head *Node) *Node {
	if nil == head {
		return nil
	}
	// a - a'
	for n := head; n != nil; n = n.Next.Next {
		n.Next = &Node{Val: n.Val, Next: n.Next}
	}
	// handler random
	for n := head; n != nil; n = n.Next.Next {
		if n.Random != nil {
			n.Next.Random = n.Random.Next
		}
	}

	// a'
	headNew := head.Next
	for node := head; node != nil; node = node.Next {
		nodeNew := node.Next
		// 重新设计旧的链表
		node.Next = node.Next.Next
		// 如果新链表后面还有数据，重新设计新链表
		if nodeNew.Next != nil {
			nodeNew.Next = nodeNew.Next.Next
		}
	}
	return headNew
}
