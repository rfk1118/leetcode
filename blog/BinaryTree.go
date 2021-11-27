package blog

type BinaryTree struct {
	key       int
	value     int
	leftNode  *BinaryTree
	rightNode *BinaryTree
}

func newBinaryTree(key, value int, left, right *BinaryTree) *BinaryTree {
	return &BinaryTree{
		key:       key,
		value:     value,
		leftNode:  left,
		rightNode: right,
	}
}
