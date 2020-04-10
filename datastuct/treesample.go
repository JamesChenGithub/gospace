package main

import (
	"./gotree"
	"fmt"
	"math/rand"
	"time"
)

func generate(start *int, maxchild int, maxdeep int) []gotree.GoTreeNode {
	if maxdeep == 0 {
		return nil
	}
	count := rand.Intn(maxchild)+1
	rootchilds := make([]gotree.GoTreeNode, count)
	for i:=0; i < len(rootchilds) ; i++  {
		rootchilds[i].Data = *start
		*start++
		rootchilds[i].Childs = generate(start, maxchild, maxdeep-1)
	}
	return  rootchilds
}

func main() {

	rand.Seed(time.Now().Unix())
	rootIndex := 0
	fmt.Println("创建根结点")
	root := gotree.GoTreeNode{}
	root.Data = rootIndex
	rootIndex++
	root.Childs = generate(&rootIndex, 5, 3)
	gotree.PrintGoTree(root)
}
