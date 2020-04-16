package main

import (
	"../algorithm/sort"
	"fmt"
)

func findPivot(sortedArray []int, target int) int {
	len := len(sortedArray)
	if len > 1 {
		if sortedArray[len-1]-target > target-sortedArray[0] {
			for i := 0; i < len-1; i++ {
				if sortedArray[i] < target && sortedArray[i+1] >= target {
					return i+1
				}
			}
		} else {
			for i := len - 1; i > 0; i-- {
				if sortedArray[i] >= target && sortedArray[i-1] < target {
					return i
				}
			}
		}
	}
	return len-1
}

func findSumOf(nums []int, target int) [][2]int {
	if len(nums) < 2 {
		return nil
	}

	array := make([]int, len(nums))
	copy(array, nums)

	sort.QuickSort(array, sort.Less)
	sort.PrintArray(array, "排序后的数组为：")
	pivot := len(array)
	if array[0] < 0 {
		pivot = findPivot(array, target - array[0])
	} else {
		pivot = findPivot(array, target)
	}

	retVals := make([][2]int, 0)

	i := 0
	j := pivot
	for i < j {
		if array[i] + array[j] > target {
			j--
		} else if array[i] + array[j] < target {
			i++
		} else {
			numbers := [2]int{array[i], array[j]}
			i++
			j--
			retVals = append(retVals, numbers)
		}
	}

	return retVals
}
func main() {
	nums := []int{1, 2, 3, 5, 6, 7, -1, -2, 8, 12, 11, 10, 9, 89}
	{
		retVals := findSumOf(nums, 7)
		fmt.Println("查找结果为:", retVals)
	}

	{
		retVals := findSumOf(nums, 10)
		fmt.Println("查找结果为:", retVals)
	}

	{
		retVals := findSumOf(nums, -10)
		fmt.Println("查找结果为:", retVals)
	}

	{
		retVals := findSumOf(nums, 6)
		fmt.Println("查找结果为:", retVals)
	}
}
