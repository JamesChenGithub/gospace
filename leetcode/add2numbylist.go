package main

import (
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
	for next != nil  {
		if next.Value.(int) < 0 {
			negative = true
			break
		}
		next = next.Next
	}

	return negative
}

func addTwoNumbers(l1 gostruct.GoLinkList, l2 gostruct.GoLinkList) *gostruct.GoLinkList {
	if l1.Empty()  {
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
		for n1 != nil || n2 != nil  {
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
			if value > 10 {
				value = value % 10
				hasEnterOne = 1
				resultList.Append(value)
			} else {
				resultList.Append(value)
				hasEnterOne = 0
			}
		}

		if hasEnterOne > 0 {
			resultList.Append(hasEnterOne)
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

	for num > 0  {
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

func main() {
	{
		num1 := 9234
		num2 := 2349
		l1 := convertIntToList(num1)
		l2 := convertIntToList(num2)
		sum := addTwoNumbers(*l1, *l2)
		sum.Print()
		fmt.Println(convertListToNum(*sum))
	}
}
