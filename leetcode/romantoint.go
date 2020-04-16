package main

import (
	"errors"
	"fmt"
)
//重复数次：一个罗马数字重复几次，就表示这个数的几倍。
//右加左减：
//在较大的罗马数字的右边记上较小的罗马数字，表示大数字加小数字。
//在较大的罗马数字的左边记上较小的罗马数字，表示大数字减小数字。
//左减的数字有限制，仅限于I、X、C。比如45不可以写成VL，只能是XLV
//但是，左减时不可跨越一个位值。比如，99不可以用IC（{\displaystyle 100-1}100-1）表示，而是用XCIX（{\displaystyle [100-10]+[10-1]}[100-10]+[10-1]）表示。（等同于阿拉伯数字每位数字分别表示。）
//左减数字必须为一位，比如8写成VIII，而非IIX。
//右加数字不可连续超过三位，比如14写成XIV，而非XIIII。（见下方“数码限制”一项。）
//同一数码最多只能连续出现三次，如40不可表示为XXXX，而要表示为XL。

func romanToInt(s string) (int, error) {
	romanMap := map[uint8]uint16{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sLen := len(s)
	valArray := make([]uint16, sLen)
	sum := 0
	sameStart := 0

	for i := 0; i < sLen; i++ {
		var ch uint8 = s[i]
		v, ok := romanMap[ch]
		if !ok {
			// not a right roman string
			return 0, errors.New("not a right roman string")
		} else {
			valArray[i] = v
		}
		sum = sum + int(v)

		if i > 0 {
			if valArray[i] > valArray[i-1] {
				var vv int = int(valArray[i-1] << 1)
				sum = sum - vv
				valArray[i-1] = -valArray[i-1]
				sameStart = i
			} else if valArray[i] == valArray[i-1]  {
				if i - sameStart >= 3 {
					return 0, errors.New("not a right roman string")
				}

			}  else {
				sameStart = i
			}
		}
	}

	return sum, nil
}
func main() {

	romanFunc := func(s string) {

		v, err := romanToInt(s)

		fmt.Println(s , " : ", v, err)
	}
	romanFunc("IIII")
	romanFunc("IV")
	romanFunc("VIII")
	romanFunc("VIIII")
	romanFunc("IX")
	romanFunc("XIII")
	romanFunc("XV")
	romanFunc("XIV")
	romanFunc("I")
	romanFunc("II")
	romanFunc("III")

}
