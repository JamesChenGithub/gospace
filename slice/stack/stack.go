package main

import "fmt"

type stack struct {
	iStack []interface{}
}

func (s *stack)push(param interface{})  {
	if param != nil {
		s.iStack = append(s.iStack, param)
	}
}

func (s stack)top() interface{}  {
	len := len(s.iStack)
	if len > 0 {
		top := s.iStack[len - 1]
		return top
	}
	return nil
}

func (s *stack)pop() interface{} {
	thetop := s.top()
	if thetop != nil {
		len :=len(s.iStack)
		s.iStack = s.iStack[0:len-1]
	}
	return thetop
}

func main() {
	istack := new(stack)

	istack.push(1)
	istack.push(2)
	istack.push(3)


	fmt.Printf("pop : %d\n", istack.pop())
	fmt.Printf("pop : %d\n", istack.pop())
	fmt.Printf("pop : %d\n", istack.pop())
}
