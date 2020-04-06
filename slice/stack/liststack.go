package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	list := list.New()
	return &Stack{list}
}

func (stack *Stack) push(value interface{}) {
	stack.list.PushBack(value)
}

func (stack *Stack) pop() interface{} {
	top := stack.list.Back()
	if top != nil {
		stack.list.Remove(top)
		return top.Value
	}
	return nil
}

func (stack Stack) top() interface{} {
	top := stack.list.Back()
	if top != nil {
		return top.Value
	}
	return nil
}

func (stack Stack) Len() int {
	return stack.list.Len()
}

func (stack Stack) Empty() bool {
	return stack.list.Len() == 0
}

func main() {
	stack := NewStack()

	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.push(4)

	topv := stack.pop().(int)
	fmt.Printf("top value : %d\n", topv)

	topv = stack.pop().(int)
	fmt.Printf("top value : %d\n", topv)

	topv = stack.pop().(int)
	fmt.Printf("top value : %d\n", topv)

	topv = stack.pop().(int)
	fmt.Printf("top value : %d\n", topv)

	if !stack.Empty() {
		topv = stack.pop().(int)
		fmt.Printf("top value : %d\n", topv)
	}


}
