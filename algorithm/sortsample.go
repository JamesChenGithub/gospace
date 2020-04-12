package main

import (
	"./sort"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	{
		compareFunc := sort.Less
		fmt.Println("用例:")
		rand.Seed(time.Now().Unix())
		maxCount := rand.Intn(1024) + 1024
		array := make([]int, maxCount)//[]int{1, 3, 5, 7, 0, 2, 4, 6, 9, 8}
		for index := range array {
			array[index] = rand.Intn(1024*10)
		}

		//array := []int{1, 5, 0 , 7, 3}
		{
			copyArray := make([]int, len(array))
			copy(copyArray, array)
			sort.ViolenceSort(copyArray, compareFunc)
		}
		{
			copyArray := make([]int, len(array))
			copy(copyArray, array)
			sort.BubbleSort(copyArray, compareFunc)
		}

		{
			copyArray := make([]int, len(array))
			copy(copyArray, array)
			sort.InsertSort(copyArray, compareFunc)
		}
		{
			copyArray := make([]int, len(array))
			copy(copyArray, array)
			sort.InsertSort2(copyArray, compareFunc)
		}

		{
			copyArray := make([]int, len(array))
			copy(copyArray, array)
			sort.SelectSort(copyArray, compareFunc)
		}

		{
			copyArray := make([]int, len(array))
			copy(copyArray, array)
			sort.ShellSort(copyArray, compareFunc)
		}
	}


	//{
	//	fmt.Println("用例:")
	//	array := []int{1,2,3,4,5}
	//	{
	//		copyArray := make([]int, len(array))
	//		copy(copyArray, array)
	//		sort.ViolenceSort(copyArray, nil)
	//	}
	//	{
	//		copyArray := make([]int, len(array))
	//		copy(copyArray, array)
	//		sort.BubbleSort(copyArray, nil)
	//	}
	//}

}
