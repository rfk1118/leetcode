package linked_list

// 430. 扁平化多级双向链表
// https://leetcode-cn.com/problems/flatten-a-multilevel-doubly-linked-list/
type flattenNode struct {
	Val               int
	Prev, Next, Child *flattenNode
}

func flatten(root *flattenNode) *flattenNode {
	if nil == root {
		return root
	}
	var ans []*flattenNode
	var bfs func(node *flattenNode)
	bfs = func(node *flattenNode) {
		if nil == node {
			return
		}
		ans = append(ans, node)
		bfs(node.Child)
		node.Child = nil
		bfs(node.Next)
	}
	bfs(root)
	var dump = &flattenNode{}
	prev := dump
	for _, fn := range ans {
		fn.Prev = prev
		prev.Next = fn
		prev = prev.Next
	}
	res := dump.Next
	res.Prev = nil
	return res
}

func flattenV2(root *flattenNode) *flattenNode {
	var dfs func(node *flattenNode) (last *flattenNode)
	dfs = func(node *flattenNode) (last *flattenNode) {
		cur := node
		for cur != nil {
			// 没有孩子，保存当前图标一直向下走
			if cur.Child == nil {
				last = cur
				cur = cur.Next
			} else {
				// 这就是next
				next := cur.Next
				// 如果有孩子，则需要返回孩子的最后一个节点
				childLast := dfs(cur.Child)

				// 当前的后面节点为孩子
				cur.Next = cur.Child
				// 孩子的前置节点为当前
				cur.Child.Prev = cur

				// next连接到孩子后面
				childLast.Next = next
				if nil != next {
					// next签名是孩子节点
					next.Prev = childLast
				}
				// 孩子节点设定为空
				cur.Child = nil
				// 返回孩子节点的最后一个节点
				last = childLast
				// 向下处理
				cur = childLast.Next
			}
		}
		return
	}
	dfs(root)
	return root
}
