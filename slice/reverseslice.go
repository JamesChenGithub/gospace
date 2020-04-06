package main

import "fmt"

func reverse(list []interface{})  {
	len := len(list)
	for i:= 0; i < len/2 ; i++  {
		tmp := list[i]
		list[i] = list[len - 1 - i]
		list[len - 1 - i] = tmp
	}
}

func main() {
	array := []int{1, 2, 3}
	array = append(array, 4)
	
	alist:= make([]interface{}, len(array))
	for i := range array {
		alist[i] = array[i]
	}

	reverse(alist)
	
	fmt.Printf("%T" , alist)
	fmt.Println(alist)

	{
		ap := [5]int{3:2}
		at := [5]int{}
		fmt.Printf("%T : %v\n" , ap, ap)
		fmt.Printf("%T : %v\n" , at, at)

		a :=[10] int{1,2,3,4,5,6,7,8,9,10}//数组
		b := a[0:3]
		c := a[3:5]
		d := a[3:]
		e := d[2:3]
		fmt.Printf("%T : %v\n" , a, a)
		fmt.Printf("%T : %v\n" , b, b)
		fmt.Printf("%T : %v\n" , c, c)
		fmt.Printf("%T : %v\n" , d, d)
		fmt.Printf("%T : %v\n" , e, e)
	}
}
