package gostack

import "errors"

type GoStackArray struct {
	iStack []interface{}
}

func (stack *GoStackArray)Push(value interface{})  {
	stack.iStack = append(stack.iStack, value)
}

func (stack GoStackArray)Top() (interface{}, error) {
	slen := len(stack.iStack)
	if slen == 0 {
		return nil, errors.New("stack is empty")
	}
	return stack.iStack[slen - 1], nil
}

func (stack GoStackArray)Empty() bool  {
	slen := len(stack.iStack)
	return slen == 0
}

func (stack *GoStackArray)Pop() (interface{}, error)  {
	slen := len(stack.iStack)
	if slen == 0 {
		return nil, errors.New("stack is empty")
	}
	val := stack.iStack[slen - 1]
	stack.iStack = stack.iStack[:slen - 1]
	return val, nil
}
