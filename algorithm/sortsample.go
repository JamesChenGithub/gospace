package main

import (
	"./sort"
	"fmt"
)

func main() {
	{
		compareFunc := sort.Great
		fmt.Println("用例:")
		array := []int{1,2,3,4}
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
