package main

import (
	"./sort"
	"fmt"
)

func main() {
	{
		fmt.Println("用例:")
		array := []int{1, 3, 5, 7, 0, 2, 4, 6, 9, 8}
		//array := []int{1, 5, 0 , 7, 3}
		{
			copyArray := make([]int, len(array))
			copy(copyArray, array)
			sort.ViolenceSort(copyArray, nil)
		}
		{
			copyArray := make([]int, len(array))
			copy(copyArray, array)
			sort.BubbleSort(copyArray, nil)
		}

		{
			copyArray := make([]int, len(array))
			copy(copyArray, array)
			sort.InsertSort(copyArray, nil)
		}
		{
			copyArray := make([]int, len(array))
			copy(copyArray, array)
			sort.InsertSort2(copyArray, nil)
		}

		{
			copyArray := make([]int, len(array))
			copy(copyArray, array)
			sort.SelectSort(copyArray, nil)
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
