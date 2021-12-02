package linked_list

// 641. 设计循环双端队列
// https://leetcode-cn.com/problems/design-circular-deque/
type MyCircularDeque struct {
	Size, Max  int
	Head, Tail *MyCircularDequeNode
}

type MyCircularDequeNode struct {
	Val        int
	Next, Prev *MyCircularDequeNode
}

func newMyCircularDequeNode(val int) *MyCircularDequeNode {
	return &MyCircularDequeNode{Val: val}
}

func MyCircularDequeConstructor(k int) MyCircularDeque {
	nl := &MyCircularDequeNode{}
	nl.Prev = nl
	nl.Next = nl
	return MyCircularDeque{
		Size: 0,
		Max:  k,
		Head: nl,
		Tail: nl,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	newNode := newMyCircularDequeNode(value)
	next := this.Head.Next
	newNode.Prev = this.Head
	this.Head.Next = newNode

	next.Prev = newNode
	newNode.Next = next
	this.Size++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	newNode := newMyCircularDequeNode(value)
	prev := this.Tail.Prev

	prev.Next = newNode
	newNode.Prev = prev

	newNode.Next = this.Tail
	this.Tail.Prev = newNode
	this.Size++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.Head.Next = this.Head.Next.Next
	this.Head.Next.Prev = this.Head
	this.Size--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.Tail.Prev = this.Tail.Prev.Prev
	this.Tail.Prev.Next = this.Tail
	this.Size--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Head.Next.Val
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Tail.Prev.Val
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.Size <= 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.Size >= this.Max
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
