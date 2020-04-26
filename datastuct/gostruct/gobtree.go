package gostruct

import (
	"fmt"
)

type BTreeNode struct {
	Data  interface{}
	Left  *BTreeNode
	Right *BTreeNode
}

func (root BTreeNode) Height() int {
	if root.Left == nil && root.Right == nil {
		return 1
	}

	leftHeight := 0
	if root.Left != nil {
		leftHeight = root.Left.Height()
	}

	rightHeight := 0
	if root.Right != nil {
		rightHeight = root.Right.Height()
	}

	max := leftHeight
	if rightHeight > max {
		max = rightHeight
	}

	return max + 1
}

//=========================
func (root BTreeNode) PrintBTree() {
	root.printBTree(0)
}

func (root BTreeNode) PreOrder() {
	fmt.Print("先序遍历: [")
	root.preOrderPrint()
	fmt.Println("]")
}

func (root BTreeNode) preOrderPrint() {
	fmt.Print(root.Data, ", ")
	if root.Left != nil {
		root.Left.preOrderPrint()
	}

	if root.Right != nil {
		root.Right.preOrderPrint()
	}
}
func (root BTreeNode) PreOrder2() {
	fmt.Print("先序遍历2: [")
	root.preOrderPrint3()
	fmt.Println("]")
}

func (root BTreeNode) preOrderPrint2() {

	stack := GoStackArray{}
	node := &root
	for node != nil || !stack.Empty() {
		if node != nil {
			fmt.Print(node.Data, ", ")
			stack.Push(node)
			node = node.Left
		} else {
			data, ok := stack.Pop()
			if ok == nil && data != nil {
				node = data.(*BTreeNode).Right
			}
		}
	}
}

func (root BTreeNode) preOrderPrint3() {

	stack := GoStackArray{}
	node := &root
	for node != nil || !stack.Empty() {

		fmt.Print(node.Data, ", ")

		if node.Right != nil {
			stack.Push(node.Right)
		}

		if node.Left != nil {
			node = node.Left
		} else {
			data, ok := stack.Pop()
			if ok == nil && data != nil {
				node = data.(*BTreeNode)
			} else {
				node = nil
			}
		}
	}
}

//=========================

func (root BTreeNode) InOrder() {
	fmt.Print("中序遍历: [")
	root.inOrderPrint()
	fmt.Println("]")
}

func (root BTreeNode) inOrderPrint() {

	if root.Left != nil {
		root.Left.inOrderPrint()
	}
	fmt.Print(root.Data, ", ")
	if root.Right != nil {
		root.Right.inOrderPrint()
	}
}

func (root BTreeNode) InOrder2() {
	fmt.Print("中序遍历2: [")
	root.inOrderPrint3()
	fmt.Println("]")
}

func (root BTreeNode) inOrderPrint2() {
	stack := GoStackArray{}
	node := &root
	for node != nil || !stack.Empty() {
		if node != nil {
			stack.Push(node)
			node = node.Left
		} else {
			data, ok := stack.Pop()
			if ok == nil && data != nil {
				node = data.(*BTreeNode)
				fmt.Print(node.Data, ", ")
				node = node.Right
			} else {
				node = nil
			}
		}
	}

}

func (root BTreeNode) inOrderPrint3() {
	stack := GoStackArray{}
	node := &root
	for node != nil || !stack.Empty() {
		for node != nil {
			stack.Push(node)
			node = node.Left
		}

		if !stack.Empty() {
			data, ok := stack.Pop()
			if ok == nil && data != nil {
				node = data.(*BTreeNode)
				fmt.Print(node.Data, ", ")
				node = node.Right
			} else {
				node = nil
			}

		}

	}

}

//===============
func (root BTreeNode) PostOrder() {
	fmt.Print("后序遍历: [")
	root.postOrderPrint()
	fmt.Println("]")
}
func (root BTreeNode) postOrderPrint() {

	if root.Left != nil {
		root.Left.postOrderPrint()
	}
	if root.Right != nil {
		root.Right.postOrderPrint()
	}
	fmt.Print(root.Data, ", ")
}

func (root BTreeNode) PostOrder2() {
	fmt.Print("后序遍历2: [")
	root.postOrderPrint2()
	fmt.Println("]")
}
func (root BTreeNode) postOrderPrint2() {
	stack := GoStackArray{}
	node := &root
	for node != nil {
		stack.Push(node)
		node = node.Left
	}

	var pLastNode *BTreeNode = nil
	for !stack.Empty() {
		data, err := stack.Pop()
		if err == nil && data != nil {
			pCur := data.(*BTreeNode)

			if pCur.Right == nil || pCur.Right == pLastNode {
				fmt.Print(pCur.Data, ", ")
				pLastNode = pCur

			} else {
				stack.Push(pCur)
				pCur = pCur.Right
				for pCur != nil  {
					stack.Push(pCur)
					pCur = pCur.Left
				}
			}
		}

	}



}

//=======================================
func (root BTreeNode) printBTree(deep int) {
	if deep < 0 {
		return
	}
	// 打印枝信息
	if deep == 0 {
		fmt.Println(root.Data)
	} else if deep == 1 {
		fmt.Printf("|-- %v\n", root.Data)
	} else if deep >= 2 {

		str := "|"
		for i := 1; i < deep; i++ {
			str += "    "
		}
		str += "|-- "
		fmt.Printf("%v%v\n", str, root.Data)

	}
	// 打印value
	//for _, child := range treeRoot.Childs {
	//	printTree(child, deep+1)
	//}
	if root.Left != nil {
		root.Left.printBTree(deep + 1)
	}

	if root.Right != nil {
		root.Right.printBTree(deep + 1)
	}

}
