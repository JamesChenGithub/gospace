package main

import (
	"./gostruct"
	"fmt"
)
func main() {
	list := gostruct.GoLinkList{}

	list.Append(1)
	list.Append(2)
	node := list.Insert(0)
	list.Insert(-1)
	list.InsertAfter(node, 3)
	list.Print()
	fmt.Println(list.Length())
	fmt.Println(list.Find(4))
	fmt.Println(list.Find(1))

	list.ForEach(func(_ int, node *gostruct.GoLinkListNode) {
		if node != nil {
			fmt.Print(node, "  ")
		}
	})
	fmt.Println()

	list.Reverse()
	list.Print()
	fmt.Println("===========below copy")
	listcopy := list.Copy()
	listcopy.Append(-2)

	list.Print()
	listcopy.Print()

	fmt.Println("===========below copy")
	listcopy.Search(func(index int, node *gostruct.GoLinkListNode) bool {
		if node.Value == 3 {
			node.Value = 5
			return true
		}
		return false
	})
	listcopy.Print()

	fmt.Println("===========check Circle")
	node = listcopy.Append(6)
	listcopy.Append(7)
	listcopy.Append(8)
	//last :=
	//last.Next = node
	listcopy.Append(9)

	fmt.Println("has circle", listcopy.HasCircle())
	if !(listcopy.HasCircle()) {
		listcopy.Print()
	}
}
