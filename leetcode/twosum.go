package main

import (
	"../algorithm/sort"
	"../datastuct/gostruct"
	"fmt"
)

func twoSumByViolence(num []int, target int) [][]int {
	if len(num) <= 1 {
		return nil
	}
	result := make([][]int, 0)
	for i := 0; i < len(num); i++  {
		for j := i+1; j < len(num); j++  {
			if  num[i] + num[j] == target{
				item := []int{num[i], num[j]}
				result = append(result, item)
			}
		}
	}
	return result
}

func twoSumBySort(num []int, target int) [][]int {
	if len(num) <= 1 {
		return nil
	}
	numCpy := gostruct.CopyIntArray(num)
	sort.InsertSort(numCpy, nil)

	result := make([][]int, 0)
	start := 0
	index := len(numCpy) - 1
	for numCpy[index] + numCpy[start] > target  {
		index--
	}

	for start < index  {
		if numCpy[index] + numCpy[start] > target {
			index--
		} else if numCpy[index] + numCpy[start] < target {
			start++
		} else {
			item := []int{numCpy[start], numCpy[index]}
			result = append(result, item)
			start++
			index--
		}
	}
	return result
}

func twoSumByMap(num []int, target int) [][]int {
	if len(num) <= 1 {
		return nil
	}
	result := make([][]int, 0)

	tarMap := make(map[int]int)

	for i := 0; i < len(num); i++  {
		key := target - num[i]
		v, ok := tarMap[key]
		if ok {
			item := []int{v, num[i]}
			result = append(result, item)
		} else {
			tarMap[num[i]] = num[i]
		}
	}

	return result
}

func main() {

	fmt.Println(twoSumByViolence([]int{}, 5))
	array := []int{1, 5, 4 , 7, 3 , 9 , 10, 2, 8, 6, 12}
	fmt.Println("twoSumByViolence=====")
	fmt.Println(twoSumByViolence(array, 5))
	fmt.Println(twoSumByViolence(array, 7))
	fmt.Println(twoSumByViolence(array, 8))
	fmt.Println(twoSumByViolence(array, 10))
	fmt.Println("twoSumByMap=====")
	fmt.Println(twoSumByMap(array, 5))
	fmt.Println(twoSumByMap(array, 7))
	fmt.Println(twoSumByMap(array, 8))
	fmt.Println(twoSumByMap(array, 10))


	fmt.Println(twoSumBySort(array, 5))
	fmt.Println(twoSumBySort(array, 7))
	fmt.Println(twoSumBySort(array, 8))


	
}
