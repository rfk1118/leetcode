package blog

import "fmt"

func preOrder(head *BinaryTree) {
	if nil == head {
		return
	}
	// 访问当前节点
	fmt.Printf("key:%d,value:%d\n", head.key, head.value)
	// 调用当前节点左节点进行遍历
	preOrder(head.leftNode)
	// 调用当前节点的右节点进行遍历
	preOrder(head.rightNode)
}

func preOrderWithLoopV1(head *BinaryTree) {
	if nil == head {
		return
	}
	stack := []*BinaryTree{head}
	for len(stack) > 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Printf("key:%d,value:%d\n", pop.key, pop.value)
		if nil != pop.rightNode {
			stack = append(stack, pop.rightNode)
		}
		if nil != pop.leftNode {
			stack = append(stack, pop.leftNode)
		}
	}
}

func preOrderWithLoopV2(head *BinaryTree) {
	var stack []*BinaryTree
	node := head
	for node != nil || len(stack) > 0 {
		for node != nil {
			fmt.Printf("key:%d,value:%d\n", node.key, node.value)
			stack = append(stack, node)
			node = node.leftNode
		}
		node = stack[len(stack)-1]
		node = node.rightNode
		stack = stack[:len(stack)-1]
	}
}

// Preorder traversal without recursion and without stack
func preOrderWithLoopMorris(root *BinaryTree) {
	if nil == root {
		return
	}
	var predecessor *BinaryTree
	var current = root
	for current != nil {
		// If left child is null, print the current node data. Move to
		// right child.
		if current.leftNode == nil {
			fmt.Printf("key:%d,value:%d\n", current.key, current.value)
			current = current.rightNode
		} else {
			// Find inorder predecessor
			predecessor = current.leftNode
			for predecessor.rightNode != nil && predecessor.rightNode != current {
				predecessor = predecessor.rightNode
			}

			// If the right child of inorder predecessor
			// already points to this node
			if predecessor.rightNode == nil {
				predecessor.rightNode = current
				current = current.leftNode
			} else {
				// If right child doesn't point to this node, then print
				// this node and make right child point to this node
				//  do Something
				fmt.Printf("key:%d,value:%d\n", current.key, current.value)
				predecessor.rightNode = nil
				current = current.rightNode
			}
		}
	}
}
