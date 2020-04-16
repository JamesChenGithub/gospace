package main

import (
	"fmt"
	"math"
)

func reverseViolence(x int) int {
	//if x < 10 && x > -10 {
	//	return x
	//}
	//
	//flag := 1
	//if x < 0 {
	//	flag = -1
	//	x = -x
	//}
	fmt.Println("reverseViolence原数 : ", x)
	retX := 0
	for x != 0 {
		if retX > math.MaxInt64/10 || retX < math.MinInt64/10 {
			fmt.Println("reverseViolence结果 : ", 0)
			return 0
		}
		retX = retX*10 + x%10
		x = x / 10
	}
	fmt.Println("reverseViolence结果 : ", retX)
	return retX
}

func reverseRecurInner(x int, deep *int) int {
	if x < 0 {
		return 0 - reverseRecurInner(-x, deep)
	}

	if x < 10 {
		return x
	}

	cur := x % 10

	newDeep := *deep + 1
	nextCur := reverseRecurInner(x/10, &newDeep)
	for i := 0; i < newDeep; i++ {
		cur = cur * 10
	}
	result := cur + nextCur
	*deep++
	return result
}

func reverseRecur2(x int, deep int) (int, bool) {
	if x/10 == 0 {
		return x, true
	}
	cur := x % 10
	for i := 0; cur != 0 && i < deep-1; i++ {
		cur = cur * 10
		if cur > math.MaxInt64/10 || cur < math.MinInt64/10 {
			return 0, false
		}
	}

	deep--
	nextCur, ok := reverseRecur2(x/10, deep)
	if !ok {
		return 0, false
	} else {
		if cur+nextCur > math.MaxInt64 || cur+nextCur < math.MinInt64 {
			return 0, false
		}
		return cur + nextCur, true
	}
}

func reverseRecur(x int) int {
	fmt.Println("reverseRecur原数 : ", x)
	value := x
	deep := 0
	for value != 0 {
		value = value / 10
		deep++
	}

	value, _ = reverseRecur2(x, deep)
	fmt.Println("reverseRecur : ", value)
	return value
}

func main() {
	//fmt.Println(reverseRecur(123))
	//fmt.Println(reverseViolence(1112312323312456789))
	//fmt.Println("暴力方式")
	//fmt.Println(reverseViolence(math.MaxInt32))
	//fmt.Println(reverseViolence(math.MaxInt64))
	//fmt.Println(reverseViolence(math.MinInt64))
	//fmt.Println(reverseViolence(math.MinInt32))
	//fmt.Println(reverseViolence(1))
	//fmt.Println(reverseViolence(123))
	//fmt.Println(reverseViolence(-321))
	//fmt.Println(reverseViolence(-120))
	//fmt.Println(reverseViolence(120))
	//fmt.Println(reverseViolence(123456789))
	//
	//fmt.Println(reverseViolence(1230123))

	compare := func(x int) bool {
		return reverseViolence(x) == reverseRecur(x)
	}
	fmt.Println(compare(1112312323312456789))
	fmt.Println(compare(1))
	fmt.Println(compare(12))
	fmt.Println(compare(123))
	fmt.Println(compare(120))
	fmt.Println(compare(201))
	fmt.Println(compare(2010))
	fmt.Println(compare(math.MaxInt32))
	fmt.Println(compare(math.MinInt32))
	fmt.Println(compare(math.MaxInt64))
	fmt.Println(compare(math.MinInt64))


	//fmt.Println("递归方式")
	//fmt.Println(reverseRecur(1) == reverseViolence(1))
	//fmt.Println(reverseRecur(12) == reverseRecur(12))
	//fmt.Println(reverseRecur(123) == reverseRecur(12))
	//fmt.Println(reverseRecur(-321) == reverseRecur(12))
	//fmt.Println(reverseRecur(-120)== reverseRecur(12))
	//fmt.Println(reverseRecur(120) == reverseRecur(12))
	//fmt.Println(reverseRecur(123456789)== reverseRecur(12))
	//fmt.Println(reverseRecur(1230123)== reverseRecur(12))
	//fmt.Println(reverseRecur(math.MaxInt32)== reverseRecur(12))
	//fmt.Println(reverseRecur(math.MaxInt64)==reverseRecur(12))
	//fmt.Println(reverseRecur(math.MinInt64)==reverseRecur(12))
	//fmt.Println(reverseRecur(math.MinInt32)==reverseRecur(12))
	//fmt.Println(reverseRecur(1112312323312456789)==reverseRecur(12))

}
