package linked_list

// 面试题 03.03. 堆盘子
// https://leetcode-cn.com/problems/stack-of-plates-lcci/
type StackOfPlates struct {
	Cap    int
	Stacks [][]int
}

func StackOfPlatesConstructor(cap int) StackOfPlates {
	return StackOfPlates{
		Cap:    cap,
		Stacks: make([][]int, 0, cap),
	}
}
func (this *StackOfPlates) Push(val int) {
	if this.Cap == 0 {
		return
	}
	stack := make([]int, 0)
	if len(this.Stacks) == 0 || len(this.Stacks[len(this.Stacks)-1]) == this.Cap {
		stack = append(stack, val)
		this.Stacks = append(this.Stacks, stack)
	} else {
		stack = this.Stacks[len(this.Stacks)-1]
		stack = append(stack, val)
		this.Stacks[len(this.Stacks)-1] = stack
	}
}

func (this *StackOfPlates) Pop() int {
	if len(this.Stacks) == 0 {
		return -1
	}
	stack := this.Stacks[len(this.Stacks)-1]
	top := -1
	if len(stack) > 0 {
		top = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
	}
	if len(stack) == 0 {
		this.Stacks = this.Stacks[0 : len(this.Stacks)-1]
	} else {
		this.Stacks[len(this.Stacks)-1] = stack
	}
	return top
}

// 实现一个popAt(int index)方法，根据指定的子栈，执行pop操作
func (this *StackOfPlates) PopAt(index int) int {
	if index >= len(this.Stacks) {
		return -1
	}
	stack := this.Stacks[index]
	top := -1
	if len(stack) > 0 {
		top = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
	}
	if len(stack) == 0 {
		if index == len(this.Stacks)-1 { // 若指定的栈是最后一个栈，直接删除
			this.Stacks = this.Stacks[0:index]
		} else { // 不是最后一个栈，使用切片特性，组成新的总栈
			this.Stacks = append(this.Stacks[0:index], this.Stacks[index+1:]...)
		}
	} else {
		this.Stacks[index] = stack
	}
	return top
}
