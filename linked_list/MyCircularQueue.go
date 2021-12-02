package linked_list

// MyCircularQueue 622. 设计循环队列
// https://leetcode-cn.com/problems/design-circular-queue/
type MyCircularQueue struct {
	Size, Max  int
	Head, Tail *MyCircularQueueNode
}

type MyCircularQueueNode struct {
	Val        int
	Next, Prev *MyCircularQueueNode
}

func newMyCircularQueueNode(val int) *MyCircularQueueNode {
	return &MyCircularQueueNode{Val: val}
}

func MyCircularQueueConstructor(k int) MyCircularQueue {
	nL := newMyCircularQueueNode(0)
	nL.Next = nL
	nL.Prev = nL
	return MyCircularQueue{Size: 0, Max: k, Head: nL, Tail: nL}
}

// 向循环队列插入一个元素。如果成功插入则返回真。
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	newNode := newMyCircularQueueNode(value)

	prev := this.Tail.Prev
	prev.Next = newNode
	newNode.Prev = prev

	this.Tail.Prev = newNode
	newNode.Next = this.Tail
	this.Size++
	return true
}

// 从循环队列中删除一个元素。如果成功删除则返回真。
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.Head.Next = this.Head.Next.Next
	this.Head.Next.Prev = this.Head
	this.Size--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Head.Next.Val
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Tail.Prev.Val
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.Size <= 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.Size >= this.Max
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
