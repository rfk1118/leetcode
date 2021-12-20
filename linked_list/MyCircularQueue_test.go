package linked_list

import (
	"fmt"
	"testing"
)

func TestMyCircularQueueConstructor(t *testing.T) {
	circularQueue := MyCircularQueueConstructor(3)
	circularQueue.EnQueue(1)
	circularQueue.EnQueue(2)
	circularQueue.EnQueue(3)
	circularQueue.EnQueue(4)
	rear := circularQueue.Rear()
	fmt.Println(rear)
	circularQueue.IsFull()
	circularQueue.DeQueue()
	circularQueue.EnQueue(4)
	circularQueue.Rear()
}
