package gostruct

import (
	"errors"
	"fmt"
)

type GoLinkListNode struct {
	Value interface{}
	Next  *GoLinkListNode
}

type GoLinkList = GoLinkListNode

func (list GoLinkList) Length() int {
	len := 0
	next := list.Next
	for next != nil {
		len++
		next = next.Next
	}

	return len
}

/**
在尾部添加
*/
func (list *GoLinkList) Append(value interface{}) *GoLinkListNode {
	next := list

	for next.Next != nil {
		next = next.Next
	}
	node := &GoLinkListNode{Value: value, Next: nil}
	next.Next = node
	return node
}

/**
在头部添加
*/
func (list *GoLinkList) Insert(value interface{}) *GoLinkListNode {
	next := list
	var nextOfNext *GoLinkListNode = nil
	if next.Next != nil {
		nextOfNext = next.Next
	}
	node := &GoLinkListNode{value, nextOfNext}
	list.Next = node
	return node
}

func (list *GoLinkList) InsertAfter(node *GoLinkListNode, value interface{}) (succ bool, listNode *GoLinkListNode, err error) {
	if node == nil {
		return false, nil, errors.New("node is nil")
	}
	next := list.Next
	hasFind := false
	for next != nil {
		if next == node {
			hasFind = true
			break
		}
		next = next.Next
	}

	if hasFind {
		var nextNext *GoLinkListNode = nil
		nextNext = next.Next
		node := &GoLinkListNode{value, nextNext}
		next.Next = node
		return true, node, nil
	} else {
		return false, nil, errors.New("node not in the list")
	}
}

func (list GoLinkList) Find(value interface{}) (bool, *GoLinkListNode) {
	next := list.Next
	for next != nil {
		if next.Value == value {
			return true, next
		}
		next = next.Next
	}

	return false, nil
}

func (list *GoLinkList) ForEach(foreach func(index int, node *GoLinkListNode)) {
	if foreach != nil {
		next := list.Next
		index := 0
		for next != nil {
			foreach(index, next)
			next = next.Next
			index++
		}
	}
}

func (list *GoLinkList) Search(seachFunc func(index int, node *GoLinkListNode) bool) {
	if seachFunc != nil {
		next := list.Next
		index := 0
		for next != nil {
			finded := seachFunc(index, next)
			if finded {
				break
			}
			next = next.Next
			index++
		}
	}
}

func (list *GoLinkListNode) Reverse() {
	next := list.Next

	if next != nil && next.Next == nil {
		return
	}
	var nextPre *GoLinkListNode = nil
	for next != nil {
		tmp := next.Next
		next.Next = nextPre
		nextPre = next
		next = tmp
	}
	list.Next = nextPre
}

func (list GoLinkList) Print() {
	next := list.Next

	for next != nil {
		fmt.Printf("(%v)", next.Value)
		if next.Next != nil {
			fmt.Print("-->")
		}
		next = next.Next
	}

	fmt.Println()
}

func (list GoLinkList) Copy() *GoLinkList {
	newList := &GoLinkList{}

	newNext := newList

	next := list.Next

	for next != nil {
		newNext.Next = &GoLinkListNode{
			Value: next.Value,
			Next:  nil,
		}
		newNext = newNext.Next
		next = next.Next
	}
	return newList
}

func (list GoLinkList)HasCircle() bool  {
	sNext := list.Next
	var fNext *GoLinkListNode = nil
	if  sNext != nil && sNext.Next != nil && sNext.Next.Next != nil {
		fNext = sNext.Next.Next
	}

	for sNext != nil && fNext != nil && sNext != fNext {
		sNext = sNext.Next
		if  fNext.Next != nil && fNext.Next.Next != nil {
			fNext = fNext.Next.Next
		} else {
			fNext = nil
		}
	}

	if fNext == nil {
		return false
	} else {
		return true
	}
}
