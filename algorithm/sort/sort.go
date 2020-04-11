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

func ViolenceSort(array []int, compare func(int, int) bool) {
	fmt.Println("======= 暴力排序:开始 =========")
	compareCount := 0
	swapCount := 0
	defer func() {
		PrintArray(array, "排序结果为")
		fmt.Println("======= 暴力排序:结束, 共比较交换:(", compareCount, ",", swapCount , ")次 =========\n\n")
	}()
	PrintArray(array, "原数组")
	if compare == nil {
		compare = func(i int, i2 int) bool {
			return i <= i2
		}
	}

	len := len(array)
	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			com := compare(array[i], array[j])
			compareCount++
			if !com {
				Swap(&array[i], &array[j])
				swapCount++
				PrintArray(array,"第" + strconv.Itoa(i)+ "-" + strconv.Itoa(compareCount) + "排序")
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
		fmt.Println("======= 暴力排序:结束, 共比较交换:(", compareCount, ",", swapCount , ")次 =========\n\n")
	}()
	PrintArray(array, "原数组")
	if compare == nil {
		compare = func(i int, i2 int) bool {
			return i <= i2
		}
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
				PrintArray(array,"第" + strconv.Itoa(i)+ "-" + strconv.Itoa(compareCount) + "排序")
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
		fmt.Println("======= 暴力排序:结束, 共比较交换:(", compareCount, ",", swapCount , ")次 =========\n\n")
	}()
	PrintArray(array, "原数组")
	if compare == nil {
		compare = func(i int, i2 int) bool {
			return i <= i2
		}
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
					PrintArray(array,"第" + strconv.Itoa(i)+ "-" + strconv.Itoa(compareCount) + "排序")
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
		fmt.Println("======= 暴力排序:结束, 共比较交换:(", compareCount, ",", swapCount , ")次 =========\n\n")
	}()
	PrintArray(array, "原数组")
	if compare == nil {
		compare = func(i int, i2 int) bool {
			return i <= i2
		}
	}

	len := len(array)

	for i := 1; i < len; i++ {
		compareCount++
		isIBig := compare(array[i-1], array[i])
		if !isIBig {
			Swap(&array[i-1], &array[i])
			swapCount++
			PrintArray(array,"第" + strconv.Itoa(i)+ "-" + strconv.Itoa(compareCount) + "排序")
			for j := i - 1; j >= 1; j-- {
				compareCount++
				isJBig := compare(array[j-1], array[j])
				if !isJBig {
					Swap(&array[j-1], &array[j])
					swapCount++
					PrintArray(array,"第" + strconv.Itoa(i)+ "-" + strconv.Itoa(compareCount) + "排序")
				}
			}
		}

		tip := "第" + strconv.Itoa(i) + "排序"
		PrintArray(array, tip)
	}
}
