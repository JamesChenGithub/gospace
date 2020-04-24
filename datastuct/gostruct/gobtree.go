package gostruct

import "fmt"

type BTreeNode struct {
	Data 	interface{}
	Left 	*BTreeNode
	Right 	*BTreeNode
}

func PrintBTree(root BTreeNode)  {
	printBTree(root, 0)
}

//=======================================
func printBTree(treeRoot BTreeNode, deep int)  {
	if deep < 0 {
		return
	}
	// 打印枝信息
	if deep == 0 {
		fmt.Println(treeRoot.Data)
	} else if deep == 1 {
		fmt.Printf("|-- %v\n", treeRoot.Data)
	} else if deep >= 2 {

		str := "|"
		for i := 1; i < deep ; i++ {
			str += "    "
		}
		str += "|-- "
		fmt.Printf("%v%v\n", str, treeRoot.Data)

	}
	// 打印value
	//for _, child := range treeRoot.Childs {
	//	printTree(child, deep+1)
	//}
	if treeRoot.Left != nil {
		printBTree(*(treeRoot.Left), deep + 1)
	}

	if treeRoot.Right != nil {
		printBTree(*(treeRoot.Right), deep + 1)
	}

}

