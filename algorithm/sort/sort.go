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
		if index != len - 1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}

func Swap(a *int, b *int)  {
	if a != b {
		t := *a
		*a = *b
		*b = t
	}
}

func ViolenceSort(array []int, compare func(int, int) bool)  {
	fmt.Println("======= 暴力排序:开始 =========")
	defer func() {
		PrintArray(array, "排序结果为")
		fmt.Println("======= 暴力排序:结束 =========\n\n")
	}()
	PrintArray(array, "原数组")
	if compare == nil {
		compare = func(i int, i2 int) bool {
			return i <= i2
		}
	}

	len := len(array)
	for i := 0; i < len; i++   {
		for j := i + 1; j < len ; j++  {
			com := compare(array[i], array[j])
			if !com {
				Swap(&array[i], &array[j])
			}
		}
		tip := "第" + strconv.Itoa(i + 1) + "排序"
		PrintArray(array, tip)
	}

}

func BubbleSort(array []int, compare func(int, int) bool)  {
	fmt.Println("======= 冒泡排序:开始 =========")
	defer func() {
		PrintArray(array, "排序结果为")
		fmt.Println("======= 冒泡排序:结束 =========\n\n")
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
	for hasSwap   {
		hasSwap = false
		lastSwapIndex := 0
		for i := 0; i < loopLen ; i++ {
			isISmall := compare(array[i], array[i+1])
			if !isISmall {
				Swap(&array[i], &array[i+1])
				hasSwap = true
				lastSwapIndex = i
				fmt.Print("[swap index : ", lastSwapIndex , "] ,")
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
