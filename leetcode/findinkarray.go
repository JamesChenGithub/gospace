package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func findByRange(kArray [][]int) []int {
	if len(kArray) == 0 {
		return nil
	}
	if len(kArray) == 1 {
		return kArray[0]
	}

	maxMinVal, minMaxVal := 0, 0
	for _, array := range kArray{
		aLen := len(array)
		if maxMinVal == 0 {
			maxMinVal = array[0]
		} else if array[0] > maxMinVal {
			maxMinVal = array[0]
		}
		if minMaxVal == 0 {
			minMaxVal = array[aLen - 1]
		} else  if array[aLen - 1] < minMaxVal {
			minMaxVal = array[aLen - 1]
		}
	}

	var samevals = findSameByVRange(kArray[0], kArray[1], minMaxVal, maxMinVal)
	for i := 2; i < len(kArray); i++  {
		samevals = findSameByVRange(samevals, kArray[i], minMaxVal, maxMinVal)
		if len(samevals) == 0 {
			return nil
		}
	}

	return samevals
}

func findRange(array []int, min int, max int) (int,int) {
	li, ri := 0, len(array)
	for idx, val := range array {
		if val >= min && li == -1 {
			li = idx
		}
		if val <= max && ri == -1  {
			ri = idx
			break;
		}
	}

	return li, ri
}

func findSameByVRange(left []int , right []int, min int, max int) []int  {
	samev := make([]int, 0)

	leftMinIdx, leftMaxIdx := findRange(left, min, max)
	rightMinIdx, rightMaxIdx := findRange(right, min, max)

	
	for i ,j := leftMinIdx, rightMinIdx; i < leftMaxIdx && j < rightMaxIdx; {
		//fmt.Printf("left[%d] = %d, right[%d] = %d\n", i, left[i], j, right[j]);
		if left[i] == right[j] {
			samev = append(samev, left[i])
			i++
			j++
		} else if left[i] < right[j] {
			i++
		} else {
			j++
		}
	}
	return samev
}

//=======================

func findSame(left []int , right []int) []int  {
	samev := make([]int, 0)
	for i ,j := 0, 0; i < len(left) && j < len(right); {
		//fmt.Printf("left[%d] = %d, right[%d] = %d\n", i, left[i], j, right[j]);
		if left[i] == right[j] {
			samev = append(samev, left[i])
			i++
			j++
		} else if left[i] < right[j] {
			i++
		} else {
			j++
		}
	}
	return samev
}

func valFind(karray [][]int )[]int {
	if len(karray) == 0 {
		return nil
	}
	if len(karray) == 1 {
		return karray[0]
	}
	
	var samevals = findSame(karray[0], karray[1])
	for i := 2; i < len(karray); i++  {
		samevals = findSame(samevals, karray[i])
		if len(samevals) == 0 {
			return nil
		}
	}

	return samevals
}

func main() {
	k := 20
	karray := make([][]int, k)



	rand.Seed(time.Now().Unix())

	sk := rand.Intn(5) + 2;
	samevs := make([]int, 0)
	for i := 0; i < sk; i++  {
		samevs = append(samevs, rand.Intn(1<<10))
	}

	for i := 0; i < k; i++ {
		array := make([]int, 0)
		len := rand.Intn(30) + 10
		maxv := 1 << 10
		for j:= 0; j < len ; j++ {
			array = append(array, rand.Intn(maxv))
		}

		array = append(array, samevs...)

		sort.Ints(array)
		karray[i] = array
		fmt.Println(array)
	}



	{
		fmt.Println("valFind 找到相同元素 in  karray : ", valFind(karray))
		fmt.Println("findByRange 找到相同元素 in  karray : ", findByRange(karray))
	}


}
