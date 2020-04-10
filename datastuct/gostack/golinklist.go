package gostack

import "errors"

type GoLinkListNode struct {
	Value interface{}
	Next *GoLinkListNode
}

type GoLinkList GoLinkListNode

func (list GoLinkList)Length() int {
	len := 0
	next := list.Next
	for next != nil  {
		len++
		next = next.Next
	}

	return len
}

/**
 在尾部添加
 */
func (list *GoLinkList)Append(value interface{}) {
	next := list.Next

	for next != nil  {
		next = next.Next
	}

	next.Next = &GoLinkListNode{value, nil}
}

/**
在头部添加
*/
func (list *GoLinkList)Insert(value interface{}) {
	next := list.Next
	var nextnext *GoLinkListNode = nil
	if next != nil {
		nextnext = next.Next
	}

	list.Next = &GoLinkListNode{value, nextnext}
}

func (list *GoLinkList)InsertAfter(node *GoLinkListNode, value interface{}) (succ bool, err error ){
	if node == nil {
		return false , errors.New("node is nil")
	}
	next := list.Next
	hasFind := false
	for next != nil  {
		if next == node {
			hasFind = true
			break
		}
		next = next.Next
	}


	if hasFind {
		var nextNext *GoLinkListNode = nil
		nextNext = next.Next
		next.Next = &GoLinkListNode{value, nextNext}
		return true,nil
	} else {
		return false , errors.New("node not in the list")
	}
}

func (list GoLinkList) Find(value interface{}) (bool, *GoLinkListNode) {
	next := list.Next
	for next != nil  {
		if next.Value == value {
			return true, next
		}
		next = next.Next
	}

	return false, nil
}


func (list *GoLinkList)ForEach(foreach func(index int, node *GoLinkListNode)) {
	if foreach != nil {
		next := list.Next
		index := 0
		for next != nil  {
			foreach(index, next)
			next = next.Next
			index++
		}
	}
}