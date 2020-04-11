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

func BubbleSort(array []int, compare func(int, int) bool)  {
	fmt.Print("======= 冒泡排序:开始 =========")
	defer fmt.Print("======= 冒泡排序:结束 =========")
	PrintArray(array, "原数组")
	if compare == nil {
		compare = func(i int, i2 int) bool {
			return i < i2
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
