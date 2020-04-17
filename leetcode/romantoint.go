package main

import (
	"../algorithm/sort"
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

func romanToInt3(s string) (int, error) {

	romanISMap := map[int]string {
		1 : "I",
		4 : "IV",
		5 : "V",
		9 : "IX",
		10 : "X",
		40 : "XL",
		50 : "L",
		90 : "XC",
		100 : "C",
		400 : "CD",
		500 : "D",
		900 : "CM",
		1000 : "M",
	}

	romanMap := map[string]int{}
	for k,v := range romanISMap {
		romanMap[v] = k
	}

	retV := 0
	for i := 0; i < len(s) ;   {
		ok := false
		kv := 0
		if i + 2 <= len(s) {
			v , k := romanMap[s[i:i+2]]
			ok = k
			kv = v
		}

		if i + 1 < len(s) && ok {
			retV += kv
			i += 2
		} else {
			retV += romanMap[string(s[i])]
			i++
		}
	}
	return retV, nil
}

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
				//左减的数字有限制，仅限于I、X、C。比如45不可以写成VL，只能是XLV
				if s[i-1] == 'I' || s[i-1] == 'X' || s[i-1] == 'C' {
					if (s[i-1] == 'I' && s[i-1] == 'C') || (s[i-1] == 'X' && s[i-1] == 'M') {
						// 左减时不可跨越一个位值
						return 0, errors.New("not a right roman string")
					} else {
						if i >= 2 && s[i-1] == s[i-2] {
							//左减数字必须为一位
							return 0, errors.New("not a right roman string")
						}
						var vv int = int(valArray[i-1] << 1)
						sum = sum - vv
						valArray[i-1] = -valArray[i-1]
						sameStart = i
					}

				} else {
					return 0, errors.New("not a right roman string")
				}

			} else if valArray[i] == valArray[i-1]  {
				if i - sameStart >= 3 {
					//同一数码最多只能连续出现三次
					//右加数字不可连续超过三位
					return 0, errors.New("not a right roman string")
				}

			}  else {
				sameStart = i
			}
		}
	}

	return sum, nil
}

// 将数字还原与罗马字符串
func intToRoman(value int) (string, error)  {
	if value <= 0 || value >= 4000 {
		return "", errors.New("超出范围")
	}

	//重复数次：一个罗马数字重复几次，就表示这个数的几倍。
	//右加左减：
	//在较大的罗马数字的右边记上较小的罗马数字，表示大数字加小数字。
	//在较大的罗马数字的左边记上较小的罗马数字，表示大数字减小数字。
	//左减的数字有限制，仅限于I、X、C。比如45不可以写成VL，只能是XLV
	//但是，左减时不可跨越一个位值。比如，99不可以用IC（{\displaystyle 100-1}100-1）表示，而是用XCIX（{\displaystyle [100-10]+[10-1]}[100-10]+[10-1]）表示。（等同于阿拉伯数字每位数字分别表示。）
	//左减数字必须为一位，比如8写成VIII，而非IIX。
	//右加数字不可连续超过三位，比如14写成XIV，而非XIIII。（见下方“数码限制”一项。）
	//同一数码最多只能连续出现三次，如40不可表示为XXXX，而要表示为XL。

	romanMap := map[int]string {
		1 : "I",
		4 : "IV",
		5 : "V",
		9 : "IX",
		10 : "X",
		40 : "XL",
		50 : "L",
		90 : "XC",
		100 : "C",
		400 : "CD",
		500 : "D",
		900 : "CM",
		1000 : "M",
	}

	allKeys := make([]int, len(romanMap))
	j := 0
	for key := range romanMap  {
		allKeys[j] = key
		j++
	}

	sort.QuickSort(allKeys, sort.Great)

	str := ""

	i := 0
	for value > 0  {

		if value - allKeys[i] >= 0  {
			str += romanMap[allKeys[i]]
			value = value - allKeys[i]
		} else {
			i++
		}
	}


	return str, nil
}

func romanToInt2(s string) int {
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
	rlt := 0
	for i := 0; i < len(s); {
		okO := false
		vO := 0
		if i+2 <= len(s) {
			v,ok :=romanMap[string(s[i:i+2])]
			vO = v
			okO = ok
		}
		if i+1 < len(s) && okO {
			rlt += vO
			i += 2
		} else {
			rlt += romanMap[string(s[i])]
			i++
		}
	}
	return rlt
}


func main() {

	testRomanFun := func(value int) {
		str, err := intToRoman(value)
		if err == nil {
			v, _ := romanToInt(str)
			v2 := romanToInt2(str)
			v3, _ := romanToInt3(str)
			if v == value && v == v2 && v == v3 {
				fmt.Println("转换成功", value , " = ", str)
				return
			}
		}
		fmt.Println("转换失败", value , " = ", str)

	}

	//testRomanFun(0)
	//testRomanFun(4000)
	testRomanFun(9)
	for i := 0; i < 4000; i++ {
		testRomanFun(i)
	}



	romanFunc := func(s string) {

		v, err := romanToInt(s)

		fmt.Println(s , " : ", v, err)
	}
	romanFunc("IIX")
	romanFunc("XIV")
	romanFunc("VL")
	romanFunc("IIII")
	romanFunc("IV")
	romanFunc("VIII")
	romanFunc("VIIII")
	romanFunc("IX")

	romanFunc("XIII")
	romanFunc("XV")

	romanFunc("I")
	romanFunc("II")
	romanFunc("III")

	fmt.Println("开始更多地测试===============")
	// 更多地测试
	romanTestFunc := func(s string, val int) {
		v, err := romanToInt(s)
		v2 := romanToInt2(s)
		v3, _ := romanToInt3(s)
		fmt.Println(s , " : ", v, err, "是否正确:" , v == val, "与其他人对比是否正确:", v == v2, "," , v == v3)
	}

	//罗马数字	数值	拉丁语
	romanTestFunc("I", 1)
	romanTestFunc("II", 2)
	romanTestFunc("III", 3)
	romanTestFunc("IV", 4)
	romanTestFunc("V", 5)
	romanTestFunc("VI", 6)
	romanTestFunc("VII", 7)
	romanTestFunc("VIII", 8)
	romanTestFunc("IX", 9)
	romanTestFunc("X", 10)



	//XI	11	ūndecim
	romanTestFunc("XI", 11)
	//XII	12	duodecim
	romanTestFunc("XII", 12)
	//XIII	13	tresdecim
	romanTestFunc("XIII", 13)
	//XIV	14	quattuordecim
	romanTestFunc("XIV", 14)
	//XV	15	quīndecim
	romanTestFunc("XV", 15)
	//XVI	16	sēdecim
	romanTestFunc("XVI", 16)
	//XVII	17	septendecim
	romanTestFunc("XVII", 17)
	//XVIII	18	octōdecim
	romanTestFunc("XVIII", 18)
	romanTestFunc("XIX", 19)
	romanTestFunc("XX", 20)
	//或
	//duodēvīgintī
	//XIX	19	novendecim
	//或
	//ūndēvīgintī
	//XX	20	vīgintī

	romanTestFunc("XXX", 30)
	romanTestFunc("XL", 40)
	romanTestFunc("L", 50)
	romanTestFunc("LX", 60)
	romanTestFunc("LXX", 70)
	romanTestFunc("LXXX", 80)
	romanTestFunc("XC", 90)
	romanTestFunc("XCIX", 99)
	//XXX	30	trīgintā
	//XL	40	quadrāgintā
	//L	50	quīnquāgintā
	//LX	60	sexāgintā
	//LXX	70	septuāgintā
	//LXXX	80	octōgintā
	//XC	90	nōnāgintā
	//XCIX	99	nōnāgintā novem
	//罗马数字	数值	拉丁语
	//C	100	centum
	romanTestFunc("C", 100)
	romanTestFunc("CI", 101)
	romanTestFunc("CII", 102)
	romanTestFunc("CXCIX", 199)
	romanTestFunc("CC", 200)
	romanTestFunc("CCC", 300)
	romanTestFunc("CD", 400)
	romanTestFunc("D", 500)
	romanTestFunc("DC", 600)
	romanTestFunc("DCC", 700)
	romanTestFunc("DCCC", 800)
	romanTestFunc("CM", 900)
	romanTestFunc("M", 1000)
	//CI	101	centum et ūnus
	//CII	102	centum et duo
	//CXCIX	199	centum nōnāgintā novem
	//CC	200	ducentī
	//CCC	300	trecentī
	//CD	400	quādringentī
	//D	500	quingentī
	//DC	600	sescentī
	//DCC	700	septingentī
	//DCCC	800	octingentī
	//CM	900	nongentī
	//M	1000	mīlle
	romanTestFunc("MCD", 1400)
	//MCD	1400
	romanTestFunc("MCDXXXVII", 1437)
	//MCDXXXVII	1437
	romanTestFunc("MD", 1500)
	romanTestFunc("MDCCC", 1800)
	romanTestFunc("MCM", 1900)
	romanTestFunc("MM", 2000)
	romanTestFunc("MMM", 3000)
	romanTestFunc("MMMCCCXXXIII", 3333)
	////MD	1500
	////MDCCC	1800
	////MCM	1900
	////MM	2000
	////MMM	3000
	////MMMCCCXXXIII	3333



}
