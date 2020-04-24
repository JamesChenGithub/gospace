package main

import (
	"./gostruct"
)

func generateTreeByArray(array []interface{}, index int) *gostruct.BTreeNode {
	len := len(array)
	if index >= len {
		 return nil
	}

	if array[index] != nil {
		node := gostruct.BTreeNode{}
		node.Data = array[index]
		if (2*index + 1) < len {
			node.Left = generateTreeByArray(array, 2*index + 1)
		}
		if (2*index + 2) < len {
			node.Right = generateTreeByArray(array, 2*index+2)
		}
		return &node
	} else {
		return nil
	}
}



func main() {
	intArray := []int{1,2,3,4,5,6, 7, 8, 1,2,3,4,5,6, 7, 8, }

	var interfaceArray = make([]interface{}, len(intArray))
	for i, d := range intArray  {
		interfaceArray[i] = d
	}

	root := generateTreeByArray(interfaceArray, 0)
	gostruct.PrintBTree(*root)
}