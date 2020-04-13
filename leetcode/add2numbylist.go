package main

import (
	"../algorithm/sort"
	"../datastuct/gostruct"
	"fmt"
	"math"
)

//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
//如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
//您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//示例：
//
//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/add-two-numbers
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func isNegativeNumList(list gostruct.GoLinkList) bool {
	next := list.Next

	negative := false
	for next != nil {
		if next.Value.(int) < 0 {
			negative = true
			break
		}
		next = next.Next
	}

	return negative
}

func isFirstGreater(l1 gostruct.GoLinkList, l2 gostruct.GoLinkList) bool {
	if l1.Length() > l2.Length() {
		return true
	} else if l1.Length() < l2.Length() {
		return false
	} else {
		l1.Reverse()
		defer l1.Reverse()

		l2.Reverse()
		defer l2.Reverse()

		n1 := l1.Next
		n2 := l2.Next

		for n1 != nil || n2 != nil {
			v1 := math.Abs(float64(n1.Value.(int)))
			v2 := math.Abs(float64(n2.Value.(int)))
			if v1 > v2 {
				return true
			} else if v1 < v2 {
				return false
			} else {
				n1 = n1.Next
				n2 = n2.Next
			}
		}

		if n2 == nil && n1 != nil {
			return true
		} else {
			return false
		}


	}

}

func decTwoNumbers(absGreater gostruct.GoLinkList, absLess gostruct.GoLinkList) *gostruct.GoLinkList {
	isGNe := isNegativeNumList(absGreater)
	flag := 1
	if isGNe {
		flag = -1
	}

	gl := absGreater.Next
	ll := absLess.Next
	resultList := gostruct.GoLinkList{}

	hasBrowerOne := 0

	for gl != nil || ll != nil {
		v1 := 0
		if gl != nil {
			v1 = gl.Value.(int)
			gl = gl.Next
		}

		v2 := 0
		if ll != nil {
			v2 = ll.Value.(int)
			ll = ll.Next
		}

		value := v1 + v2 + hasBrowerOne
		if flag == -1 {
			if value >= 0 {
				value += 10
				hasBrowerOne = -1
				resultList.Append(value)
			} else {
				hasBrowerOne = 0
				resultList.Append(value)
			}
		} else {
			if value >= 0 {
				hasBrowerOne = 0
				resultList.Append(value * flag)
			} else {
				value += 10
				hasBrowerOne = -1
				resultList.Append(value * flag)
			}
		}

	}

	if hasBrowerOne < 0 {
		resultList.Append(hasBrowerOne * flag)
	}

	resultList.Reverse()
	next := resultList.Next

	for next != nil {
		if next.Value.(int) == 0 && next.Next != nil {
			tmp := next.Next
			resultList.Remove(next)
			next = tmp
		} else {
			break
		}
	}

	resultList.Reverse()
	return &resultList
}

func addTwoNumbers(l1 gostruct.GoLinkList, l2 gostruct.GoLinkList) *gostruct.GoLinkList {
	if l1.Empty() {
		return l2.Copy()
	}

	if l2.Empty() {
		return l1.Copy()
	}

	n1 := l1.Next
	n2 := l2.Next

	isN1Negative := isNegativeNumList(l1)
	isN2Negative := isNegativeNumList(l2)

	resultList := gostruct.GoLinkList{}
	if isN1Negative == isN2Negative {
		// 按拉加
		hasEnterOne := 0
		for n1 != nil || n2 != nil {
			v1 := 0
			if n1 != nil {
				v1 = n1.Value.(int)
				n1 = n1.Next
			}

			v2 := 0
			if n2 != nil {
				v2 = n2.Value.(int)
				n2 = n2.Next
			}

			value := v1 + v2 + hasEnterOne
			if value >= 10 {
				value = value % 10
				hasEnterOne = 1
				resultList.Append(value)
			} else if value <= -10 {
				value = value % -10
				hasEnterOne = -1
				resultList.Append(value)
			} else {
				resultList.Append(value)
				hasEnterOne = 0
			}
		}

		if hasEnterOne != 0 {
			resultList.Append(hasEnterOne)
		}
	} else {
		firstBig := isFirstGreater(l1, l2)
		if firstBig {
			return decTwoNumbers(l1, l2)
		} else {
			return decTwoNumbers(l2, l1)
		}

	}

	return &resultList

}

func convertIntToList(num int) *gostruct.GoLinkList {
	list := gostruct.GoLinkList{}
	isNav := false
	if num < 0 {
		isNav = true
		num = (int)(math.Abs(float64(num)))
	}

	for num > 0 {
		val := num % 10
		if isNav {
			val = -val
		}
		list.Append(val)
		num = num / 10
	}

	list.Print()
	return &list
}

func convertListToNum(list gostruct.GoLinkList) int64 {
	var num int64 = 0
	index := 1
	next := list.Next
	for next != nil {
		num += int64(next.Value.(int) * index)
		index *= 10
		next = next.Next
	}

	return num

}

func add2Numbers(a int, b int) int64 {
	stats := sort.GoStats{}
	stats.StartStats()
	l1 := convertIntToList(a)
	l2 := convertIntToList(b)
	sum := addTwoNumbers(*l1, *l2)
	sum.Print()
	retv := convertListToNum(*sum)
	right := false
	result := a + b
	if int64(result) == retv {
		right = true
	}

	stats.StopStats()
	stats.PrintResult()
	fmt.Println("执行结果为：", retv, "正确与否", right)
	return retv
}

func main() {
	add2Numbers(-1232, 4291)
	add2Numbers(-1, -9)
	add2Numbers(-1, 2)
	add2Numbers(-1, -1)

	add2Numbers(-1, -19)
	add2Numbers(1, 9)
	add2Numbers(-1, 9)

	add2Numbers(-1232, 291)
	add2Numbers(1232, 291)
	add2Numbers(1232, 9291)

	add2Numbers(-5555, 5555)
	add2Numbers(4555, 5555)
	add2Numbers(0, 5555)
	add2Numbers(0, -5555)

}

// 衍生多个

// 支持加减乘除

// 衍生加减乘除法

// 衍生浮点数加减法
