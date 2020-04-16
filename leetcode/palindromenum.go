package main

import (
	"fmt"
	"strconv"
)

func isPalindromeStack(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	xLen := len(strconv.Itoa(x))
	xStack := make([]int, xLen)
	index := xLen / 2

	i := 0
	for ; i < index; i++ {
		xStack[i] = x % 10
		x = x / 10
	}
	if xLen%2 != 0 {
		x = x / 10
	}
	i--
	for j := index; j < xLen && i >= 0; j++ {
		if xStack[i] != x%10 {
			return false
		}
		i--
		x = x / 10
	}
	return x == 0
}

func isPalindromeStack2(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	xLen := len(strconv.Itoa(x))
	xStack := make([]int, xLen)
	index := xLen / 2

	for i := 0; i < xLen; i++ {
		if i < index {
			xStack[i] = x % 10
			x = x / 10
		} else {
			if i == index && xLen%2 != 0 {
				x = x / 10
				continue
			}
			if x%10 != xStack[xLen-1-i] {
				return false
			}
			x = x / 10
		}
	}

	return x == 0
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	array := make([]int, 0)

	for ; x != 0; x = x / 10 {
		array = append(array, x%10)
	}

	for i, j := 0, len(array)-1; array[i] == array[j]; i++ {
		j--
		if i >= j {
			return true
		}
	}

	return false
}

func isPalindromeString(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	str := strconv.Itoa(x)

	for i, j:= 0, len(str)-1; str[i] == str[j] ; i++  {
		j--
		if i >= j {
			return true
		}
	}

	return false
}
func main() {

	printFunc := func(x int) {
		ok := isPalindrome(x)
		ok2 := isPalindromeStack(x)
		ok3 := isPalindromeStack2(x)
		ok4 := isPalindromeString(x)
		if ok {
			fmt.Println(x, " 是回文数 : ", ok, ok2, ok3, ok4)
		} else {
			fmt.Println(x, " 不是回文数 : ", ok, ok2, ok3, ok4)
		}

	}
	printFunc(11)
	printFunc(1111)
	printFunc(121)
	printFunc(1121)
	printFunc(-1121)

	printFunc(11211)
	printFunc(112211)
	printFunc(11234211)
	printFunc(112343211)
	printFunc(1)
}
