package linked_list

// MyLinkedList 0-index
// 707. 设计链表
// https://leetcode-cn.com/problems/design-linked-list/
type MyLinkedList struct {
	Size          int
	Head, Tail, N *MyLinkedListNode
}

type MyLinkedListNode struct {
	Val        int
	Next, Prev *MyLinkedListNode
}

func newMyLinkedListNode(val int) *MyLinkedListNode {
	return &MyLinkedListNode{Val: val}
}

// Constructor MyLinkedListConstructor
func MyLinkedListConstructor() MyLinkedList {
	Nl := newMyLinkedListNode(0)
	Nl.Prev = Nl
	Nl.Next = Nl
	return MyLinkedList{
		Size: 0,
		Head: Nl,
		Tail: Nl,
		N:    Nl,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index > this.Size || this.Size == 0 {
		return -1
	}
	get := this.getNode(index)
	if get == this.N {
		return -1
	}
	return get.Val
}

func (this *MyLinkedList) getNode(index int) *MyLinkedListNode {
	c := this.Head.Next
	for i := 0; i < index && c != nil; i++ {
		c = c.Next
	}
	return c
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := newMyLinkedListNode(val)
	next := this.Head.Next

	next.Prev = newNode
	newNode.Next = next

	newNode.Prev = this.Head
	this.Head.Next = newNode

	this.Size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := newMyLinkedListNode(val)

	prev := this.Tail.Prev
	prev.Next = newNode
	newNode.Prev = prev

	this.Tail.Prev = newNode
	newNode.Next = this.Tail

	this.Size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		this.AddAtTail(val)
	}
	if this.Size < index {
		return
	}
	node := this.getNode(index)
	prev := node.Prev
	newNode := newMyLinkedListNode(val)

	prev.Next = newNode
	node.Prev = newNode
	newNode.Prev = prev
	newNode.Next = node

	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this.Size < index {
		return
	}
	node := this.getNode(index)
	if node == this.N {
		return
	}
	prev := node.Prev
	next := node.Next
	prev.Next = next
	next.Prev = prev
	this.Size--
}
