package gostruct

import "fmt"

type GoTreeNode struct {
	Data interface{}
	Childs []GoTreeNode
}

func PrintGoTree(treeRoot GoTreeNode)  {
	printTree(treeRoot, 0)
}


//=======================================
func printTree(treeRoot GoTreeNode, deep int)  {
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
			str += "   "
		}
		str += "|-- "
		fmt.Printf("%v%v\n", str, treeRoot.Data)

	}
	// 打印value

	for _, child := range treeRoot.Childs {
		printTree(child, deep+1)
	}
}