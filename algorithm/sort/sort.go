package sort

import (
	"fmt"
	"math/rand"
)

func PrintArray(array []int, tip string) {
	fmt.Print(tip, " : [")
	len := len(array)
	for index, value := range array {
		fmt.Print(value)
		if index != len-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")

	//fmt.Println(tip , " : ", array)
}

func Swap(a *int, b *int) {
	if a != b {
		t := *a
		*a = *b
		*b = t
	}
}

func Less(a int, b int) bool {
	return a <= b
}

func Great(a int, b int) bool {
	return a >= b
}

func ViolenceSort(array []int, compare func(int, int) bool) {
	fmt.Println("======= 暴力排序:开始 =========")
	stats := GoStats{}
	stats.StartStats()
	defer func() {
		PrintArray(array, "排序结果为")
		stats.StopStats()
		stats.PrintResult()
		fmt.Println("======= 暴力排序:结束 =========")
	}()
	PrintArray(array, "原数组")
	if compare == nil {
		compare = Less
	}

	len := len(array)
	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			com := compare(array[i], array[j])
			stats.CompareCount++
			if !com {
				Swap(&array[i], &array[j])
				stats.SwapCount++
				//PrintArray(array, "第"+strconv.Itoa(i)+"-"+strconv.Itoa(compareCount)+"排序")
			}
		}
		//tip := "第" + strconv.Itoa(i+1) + "排序"
		//PrintArray(array, tip)
	}

}

func BubbleSort(array []int, compare func(int, int) bool) {
	fmt.Println("======= 冒泡排序:开始 =========")
	stats := GoStats{}
	stats.StartStats()
	defer func() {
		PrintArray(array, "排序结果为")
		stats.StopStats()
		stats.PrintResult()
		fmt.Println("======= 冒泡排序:结束 =========")
	}()
	PrintArray(array, "原数组")
	if compare == nil {
		compare = Less
	}

	len := len(array)
	hasSwap := true

	index := 0
	loopLen := len - 1 - index
	for hasSwap {
		hasSwap = false
		lastSwapIndex := 0
		for i := 0; i < loopLen; i++ {
			stats.CompareCount++
			isISmall := compare(array[i], array[i+1])
			if !isISmall {
				Swap(&array[i], &array[i+1])
				hasSwap = true
				lastSwapIndex = i
				stats.SwapCount++
				//fmt.Print("[swap index : ", lastSwapIndex, "] ,")
				//PrintArray(array, "第"+strconv.Itoa(i)+"-"+strconv.Itoa(stats.CompareCount)+"排序")
			}

		}
		if hasSwap {
			loopLen = lastSwapIndex + 1
		}
		if loopLen == 1 {
			hasSwap = false
		}
		//fmt.Println()
		index++
		//tip := "第" + strconv.Itoa(index) + "排序"
		//PrintArray(array, tip)
	}
}

func InsertSort(array []int, compare func(int, int) bool) {
	fmt.Println("======= 插入排序:开始 =========")
	stats := GoStats{}
	stats.StartStats()
	defer func() {
		PrintArray(array, "排序结果为")
		stats.StopStats()
		stats.PrintResult()
		fmt.Println("======= 插入排序:结束 =========")
	}()
	PrintArray(array, "原数组")
	if compare == nil {
		compare = Less
	}

	len := len(array)

	for i := 1; i < len; i++ {
		stats.CompareCount++
		isIBig := compare(array[i-1], array[i])
		if !isIBig {
			for j := 0; j < i; j++ {
				stats.CompareCount++
				isJSmall := compare(array[j], array[i])
				if !isJSmall {
					Swap(&array[j], &array[i])
					stats.SwapCount++
					//PrintArray(array, "第"+strconv.Itoa(i)+"-"+strconv.Itoa(stats.CompareCount)+"排序")
				}
			}
		}

		//tip := "第" + strconv.Itoa(i) + "排序"
		//PrintArray(array, tip)
	}
}

func InsertSort2(array []int, compare func(int, int) bool) {
	fmt.Println("======= 插入排序2:开始 =========")
	stats := GoStats{}
	stats.StartStats()
	defer func() {
		PrintArray(array, "排序结果为")
		stats.StopStats()
		stats.PrintResult()
		fmt.Println("======= 插入排序2:结束 =========")
	}()
	PrintArray(array, "原数组")
	if compare == nil {
		compare = Less
	}

	len := len(array)

	for i := 1; i < len; i++ {
		stats.CompareCount++
		isIBig := compare(array[i-1], array[i])
		if !isIBig {
			Swap(&array[i-1], &array[i])
			stats.SwapCount++
			//PrintArray(array, "第"+strconv.Itoa(i)+"-"+strconv.Itoa(stats.CompareCount)+"排序")
			for j := i - 1; j >= 1; j-- {
				stats.CompareCount++
				isJBig := compare(array[j-1], array[j])
				if !isJBig {
					Swap(&array[j-1], &array[j])
					stats.SwapCount++
					//PrintArray(array, "第"+strconv.Itoa(i)+"-"+strconv.Itoa(stats.CompareCount)+"排序")
				}
			}
		}

		//tip := "第" + strconv.Itoa(i) + "排序"
		//PrintArray(array, tip)
	}
}

func SelectSort(array []int, compare func(int, int) bool) {
	fmt.Println("======= 选择排序:开始 =========")
	stats := GoStats{}
	stats.StartStats()
	defer func() {
		PrintArray(array, "排序结果为")
		stats.StopStats()
		stats.PrintResult()
		fmt.Println("======= 选择排序:结束 =========")
	}()
	PrintArray(array, "原数组")
	if compare == nil {
		compare = Less
	}

	len := len(array)

	for i := 0; i < len; i++ {
		min := array[i]
		index := i
		for j := i + 1; j < len; j++ {
			isMin := compare(min, array[j])
			stats.CompareCount++
			if !isMin {
				min = array[j]
				index = j
			}
		}
		if i != index {
			Swap(&array[i], &array[index])
			stats.SwapCount++
		}
		//tip := "第" + strconv.Itoa(i) + "排序"
		//PrintArray(array, tip)
	}
}

func ShellSort(array []int, compare func(int, int) bool) {
	fmt.Println("======= 希尔排序:开始 =========")
	stats := GoStats{}
	stats.StartStats()
	defer func() {
		PrintArray(array, "排序结果为")
		stats.StopStats()
		stats.PrintResult()
		fmt.Println("======= 希尔排序:结束 =========")
	}()
	PrintArray(array, "原数组")
	if compare == nil {
		compare = Less
	}

	len := len(array)
	sortCount := 1

	for gap := len / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < len; i++ {
			j := i
			for j-gap >= 0 {
				isJGapLess := compare(array[j-gap], array[j])
				stats.CompareCount++
				//fmt.Printf("compare(array[%d](%d), array[%d](%d))\n", j-gap, array[j-gap], j, array[j])
				//PrintArray(array, "第"+strconv.Itoa(compareCount)+"比较")
				if !isJGapLess {
					Swap(&array[j-gap], &array[j])
					stats.SwapCount++
					//PrintArray(array, "第"+strconv.Itoa(sortCount)+"-"+strconv.Itoa(stats.CompareCount)+"排序")
				}
				j = j - gap
			}

		}
		//tip := "第" + strconv.Itoa(sortCount) + "排序"
		//PrintArray(array, tip)
		sortCount++
	}
}


//====================================================
type PartitionFunc func (array []int, compare func(int, int) bool, left int, right int, stats *GoStats) int

func partitionLeft(array []int, compare func(int, int) bool, left int, right int, stats *GoStats) int{
	pivot := left
	index := pivot + 1
	for i := index; i <= right; i++ {
		isOK := compare(array[pivot], array[i])
		stats.CompareCount++
		//fmt.Printf("compare(array[%d](%d), array[%d](%d)) : %v\n",  pivot, array[pivot], i, array[i],isOK)
		if !isOK {
			if i != index {
				Swap(&array[i], &array[index])
				stats.SwapCount++
			}
			//fmt.Printf("Swap(array[%d](%d), array[%d](%d)) : %v\n", i, array[i], index, array[index], isOK)
			index++
		}
		//PrintArray(array, "partitionLeft")
	}
	Swap(&array[pivot], &array[index - 1])
	stats.SwapCount++
	//PrintArray(array, "partitionLeft")
	return index - 1
}

func partitionRight(array []int, compare func(int, int) bool, left int, right int, stats *GoStats) int{
	pivot := right
	index := pivot - 1
	for i := index; i >= left; i-- {
		isOK := compare(array[i], array[pivot])
		stats.CompareCount++
		//fmt.Printf("compare(array[%d](%d), array[%d](%d)) : %v\n", i, array[i], pivot, array[pivot], isOK)
		if !isOK {
			if i != index {
				Swap(&array[i], &array[index])
				stats.SwapCount++
			}
			//fmt.Printf("Swap(array[%d](%d), array[%d](%d)) : %v\n", i, array[i], index, array[index], isOK)
			index--
		}
		//PrintArray(array, "partitionRight")
	}
	Swap(&array[pivot], &array[index + 1])
	stats.SwapCount++
	//PrintArray(array, "partitionRight")
	return index + 1
}

func partitionMiddle(array []int, compare func(int, int) bool, left int, right int, stats *GoStats) int{
	mid := left + (right - left) / 2
	if left != mid {
		Swap(&array[left], &array[mid])
		stats.SwapCount++
	}

	//PrintArray(array, "partitionMiddle")
	pivot := partitionLeft(array, compare, left, right, stats)
	//PrintArray(array, "partitionMiddle")
	return pivot
}

func partitionRand(array []int, compare func(int, int) bool, left int, right int, stats *GoStats) int{
	rand := left + rand.Intn(right - left + 1)
	pivot := 0
	if rand == left {
		pivot = partitionLeft(array, compare, left, right, stats)
	} else if rand == right {
		pivot = partitionRight(array, compare, left, right, stats)
	} else {
		Swap(&array[left], &array[rand])
		stats.SwapCount++
		pivot = partitionLeft(array, compare, left, right, stats)
	}
	return pivot
}

func quickSortInner(array []int, compare func(int, int) bool, left int, right int, stats *GoStats)  {
	if left < right {
		parFunc := PartitionFunc(partitionRand)
		pivot := parFunc(array, compare, left, right, stats)
		quickSortInner(array, compare, left, pivot - 1, stats)
		quickSortInner(array, compare, pivot + 1, right, stats)
	}
}

func QuickSort(array []int, compare func(int, int) bool) {
	// fmt.Println("======= 快排序:开始 =========")
	stats := GoStats{}
	stats.StartStats()
	defer func() {
		//PrintArray(array, "排序结果为")
		stats.StopStats()
		//stats.PrintResult()
		//fmt.Println("======= 快排序:结束 =========")
	}()
	//PrintArray(array, "原数组")
	if compare == nil {
		compare = Less
	}

	quickSortInner(array, compare, 0, len(array) - 1, &stats)
}

