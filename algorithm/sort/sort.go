package sort

import (
	"fmt"
	"strconv"
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
	compareCount := 0
	swapCount := 0
	defer func() {
		PrintArray(array, "排序结果为")
		fmt.Println("======= 暴力排序:结束, 共比较交换:(", compareCount, ",", swapCount, ")次 =========\n\n")
	}()
	PrintArray(array, "原数组")
	if compare == nil {
		compare = Less
	}

	len := len(array)
	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			com := compare(array[i], array[j])
			compareCount++
			if !com {
				Swap(&array[i], &array[j])
				swapCount++
				PrintArray(array, "第"+strconv.Itoa(i)+"-"+strconv.Itoa(compareCount)+"排序")
			}
		}
		tip := "第" + strconv.Itoa(i+1) + "排序"
		PrintArray(array, tip)
	}

}

func BubbleSort(array []int, compare func(int, int) bool) {
	fmt.Println("======= 冒泡排序:开始 =========")
	compareCount := 0
	swapCount := 0
	defer func() {
		PrintArray(array, "排序结果为")
		fmt.Println("======= 冒泡排序:结束, 共比较交换:(", compareCount, ",", swapCount, ")次 =========\n\n")
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
			compareCount++
			isISmall := compare(array[i], array[i+1])
			if !isISmall {
				Swap(&array[i], &array[i+1])
				hasSwap = true
				lastSwapIndex = i
				swapCount++
				//fmt.Print("[swap index : ", lastSwapIndex, "] ,")
				PrintArray(array, "第"+strconv.Itoa(i)+"-"+strconv.Itoa(compareCount)+"排序")
			}

		}
		if hasSwap {
			loopLen = lastSwapIndex + 1
		}
		if loopLen == 1 {
			hasSwap = false
		}
		fmt.Println()
		index++
		tip := "第" + strconv.Itoa(index) + "排序"
		PrintArray(array, tip)
	}
}

func InsertSort(array []int, compare func(int, int) bool) {
	fmt.Println("======= 插入排序:开始 =========")
	compareCount := 0
	swapCount := 0
	defer func() {
		PrintArray(array, "排序结果为")
		fmt.Println("======= 插入排序:结束, 共比较交换:(", compareCount, ",", swapCount, ")次 =========\n\n")
	}()
	PrintArray(array, "原数组")
	if compare == nil {
		compare = Less
	}

	len := len(array)

	for i := 1; i < len; i++ {
		compareCount++
		isIBig := compare(array[i-1], array[i])
		if !isIBig {
			for j := 0; j < i; j++ {
				compareCount++
				isJSmall := compare(array[j], array[i])
				if !isJSmall {
					Swap(&array[j], &array[i])
					swapCount++
					PrintArray(array, "第"+strconv.Itoa(i)+"-"+strconv.Itoa(compareCount)+"排序")
				}
			}
		}

		tip := "第" + strconv.Itoa(i) + "排序"
		PrintArray(array, tip)
	}
}

func InsertSort2(array []int, compare func(int, int) bool) {
	fmt.Println("======= 插入排序:开始 =========")
	compareCount := 0
	swapCount := 0
	defer func() {
		PrintArray(array, "排序结果为")
		fmt.Println("======= 插入排序:结束, 共比较交换:(", compareCount, ",", swapCount, ")次 =========\n\n")
	}()
	PrintArray(array, "原数组")
	if compare == nil {
		compare = Less
	}

	len := len(array)

	for i := 1; i < len; i++ {
		compareCount++
		isIBig := compare(array[i-1], array[i])
		if !isIBig {
			Swap(&array[i-1], &array[i])
			swapCount++
			PrintArray(array, "第"+strconv.Itoa(i)+"-"+strconv.Itoa(compareCount)+"排序")
			for j := i - 1; j >= 1; j-- {
				compareCount++
				isJBig := compare(array[j-1], array[j])
				if !isJBig {
					Swap(&array[j-1], &array[j])
					swapCount++
					PrintArray(array, "第"+strconv.Itoa(i)+"-"+strconv.Itoa(compareCount)+"排序")
				}
			}
		}

		tip := "第" + strconv.Itoa(i) + "排序"
		PrintArray(array, tip)
	}
}

func SelectSort(array []int, compare func(int, int) bool) {
	fmt.Println("======= 选择排序:开始 =========")
	compareCount := 0
	swapCount := 0
	defer func() {
		PrintArray(array, "选择结果为")
		fmt.Println("======= 选择排序:结束, 共比较交换:(", compareCount, ",", swapCount, ")次 =========\n\n")
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
			compareCount++
			if !isMin {
				min = array[j]
				index = j
			}
		}
		if i != index {
			Swap(&array[i], &array[index])
			swapCount++
		}
		tip := "第" + strconv.Itoa(i) + "排序"
		PrintArray(array, tip)
	}
}

func ShellSort(array []int, compare func(int, int) bool) {
	fmt.Println("======= 希尔排序:开始 =========")
	compareCount := 0
	swapCount := 0
	defer func() {
		PrintArray(array, "选择结果为")
		fmt.Println("======= 希尔排序:结束, 共比较交换:(", compareCount, ",", swapCount, ")次 =========\n\n")
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
				compareCount++
				//fmt.Printf("compare(array[%d](%d), array[%d](%d))\n", j-gap, array[j-gap], j, array[j])
				//PrintArray(array, "第"+strconv.Itoa(compareCount)+"比较")
				if !isJGapLess {
					Swap(&array[j-gap], &array[j])
					swapCount++
					PrintArray(array, "第"+strconv.Itoa(sortCount)+"-"+strconv.Itoa(compareCount)+"排序")
				}
				j = j - gap
			}

		}
		tip := "第" + strconv.Itoa(sortCount) + "排序"
		PrintArray(array, tip)
		sortCount++
	}
}
