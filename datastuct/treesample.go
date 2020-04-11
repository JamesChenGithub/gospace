package main

import (
	"./gostruct"
	"fmt"
	"math/rand"
	"time"
)

func generate(start *int, maxchild int, maxdeep int) []gostruct.GoTreeNode {
	if maxdeep == 0 {
		return nil
	}
	count := rand.Intn(maxchild+1)
	rootchilds := make([]gostruct.GoTreeNode, count)
	for i:=0; i < len(rootchilds) ; i++  {
		rootchilds[i].Data = *start
		value := *start * 10
		rootchilds[i].Childs = generate(&value, maxchild, maxdeep-1)
		*start++
	}
	return  rootchilds
}

func main() {

	rand.Seed(time.Now().Unix())
	rootIndex := 0
	fmt.Println("创建根结点")
	root := gostruct.GoTreeNode{}
	root.Data = rootIndex
	rootIndex++
	root.Childs = generate(&rootIndex, 4, 5)
	gostruct.PrintGoTree(root)
}
