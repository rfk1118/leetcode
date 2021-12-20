package linked_list

// 1670. 设计前中后队列
// https://leetcode-cn.com/problems/design-front-middle-back-queue/
type FrontMiddleBackQueue struct {
	Size       int
	Head, Tail *FrontMiddleBackQueueNode
}

type FrontMiddleBackQueueNode struct {
	Val        int
	Next, Prev *FrontMiddleBackQueueNode
}

func newFrontMiddleBackQueueNode(val int) *FrontMiddleBackQueueNode {
	return &FrontMiddleBackQueueNode{
		Val: val,
	}
}

func FrontMiddleBackQueueConstructor() FrontMiddleBackQueue {
	dumpNode := newFrontMiddleBackQueueNode(0)
	dumpNode.Next = dumpNode
	dumpNode.Prev = dumpNode
	return FrontMiddleBackQueue{
		Size: 0,
		Head: dumpNode,
		Tail: dumpNode,
	}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	newNode := newFrontMiddleBackQueueNode(val)
	next := this.Head.Next
	newNode.Next = next
	next.Prev = newNode
	this.Head.Next = newNode
	newNode.Prev = this.Head
	this.Size++
}

func (this *FrontMiddleBackQueue) getMiddle() *FrontMiddleBackQueueNode {
	mid := (this.Size + 1) / 2
	c := this.Head
	for i := 0; i < mid; i++ {
		c = c.Next
	}
	return c
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	if this.IsEmpty() {
		this.PushFront(val)
	} else {
		middle := this.getMiddle()
		newNode := newFrontMiddleBackQueueNode(val)
		front := this.Size % 2
		if front == 1 {
			prev := middle.Prev
			prev.Next = newNode
			newNode.Prev = prev
			newNode.Next = middle
			middle.Prev = newNode
		} else {
			next := middle.Next
			next.Prev = newNode
			newNode.Next = next
			middle.Next = newNode
			newNode.Prev = middle
		}
		this.Size++
	}

}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	newNode := newFrontMiddleBackQueueNode(val)
	prev := this.Tail.Prev
	prev.Next = newNode
	newNode.Prev = prev
	newNode.Next = this.Tail
	this.Tail.Prev = newNode
	this.Size++
}

func (this *FrontMiddleBackQueue) PopFront() int {
	if this.IsEmpty() {
		return -1
	}
	temp := this.Head.Next
	this.Head.Next = this.Head.Next.Next
	this.Head.Next.Prev = this.Head
	this.Size--
	return temp.Val
}

func (this *FrontMiddleBackQueue) IsEmpty() bool {
	return this.Size == 0
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	if this.IsEmpty() {
		return -1
	}
	middle := this.getMiddle()
	midPrev := middle.Prev
	midNext := middle.Next
	midPrev.Next = midNext
	midNext.Prev = midPrev
	this.Size--
	return middle.Val
}

func (this *FrontMiddleBackQueue) PopBack() int {
	if this.IsEmpty() {
		return -1
	}
	tmp := this.Tail.Prev
	this.Tail.Prev = this.Tail.Prev.Prev
	this.Tail.Prev.Next = this.Tail
	this.Size--
	return tmp.Val
}

// ["FrontMiddleBackQueue","pushMiddle","pushMiddle","pushMiddle","popMiddle","popMiddle","popMiddle"]
// [[],[1],[2],[3],[],[],[]]
