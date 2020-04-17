package main

import (
	"fmt"
)

func longestCommonPrefixOf2(str1 string, str2 string) string {
	l1 := len(str1)
	if l1 == 0 {
		return str2
	}

	l2 := len(str2)
	if l2 == 0 {
		return str1
	}

	l := l1
	if l2 < l1 {
		l = l2
	}

	// 暴力
	i := 0
	for ; i < l; i++  {
		if str1[i] != str2[i] {
			break
		}
	}

	if i != 0 {
		return str1[0 : i]
	}
	return ""

}

func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	str := longestCommonPrefixOf2(strs[0], strs[1])
	if len(str) > 0 {
		newStrs := append([]string{str}, strs[2:]... )
		return longestCommonPrefix1(newStrs)
	}  else {
		return  ""
	}
}

func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	index := 0
	len0 := len(strs[0])
	hasBreak := false
	for !hasBreak {

		for i:= 1; i < len(strs) ; i++  {
			if index < len0 && index < len(strs[i]) {
				if strs[i][index] != strs[0][index] {
					hasBreak = true
					break
				}
			} else {
				hasBreak = true
				break
			}
		}
		if !hasBreak {
			index++
		}

	}

	if index > 0 {
		return strs[0][0 : index-1]
	}
	return ""

}

func main() {
	testFunc := func(strs []string) {
		str := longestCommonPrefix1(strs)
		str2 := longestCommonPrefix2(strs)

		fmt.Println( str,str2 ," 是前缀 : ", strs)

	}

	testFunc([]string{"12"})
	testFunc([]string{"12", "21"})
	testFunc([]string{"12", "123", "1234"})
	testFunc([]string{"123", "123", "1234"})
	testFunc([]string{"123","a", "1234"})


}
