package main

import (
	"../algorithm/sort"
	"fmt"
)

func fib_vol(n int, stats *sort.GoStats) uint64  {
	if n <= 1 {
		return 1
	}

	return fib_vol(n-1, stats) + fib_vol(n-2, stats)
}

func fib_dp(n int, stats *sort.GoStats) uint64  {
	if n <= 1 {
		return 1
	}
	resultMemo := make([]uint64, n+1)
	resultMemo[0] = 1
	resultMemo[1] = 1
	for i := 2; i <= n ; i++ {
		resultMemo[i] = resultMemo[i-1] + resultMemo[i-2]
	}
	return resultMemo[n]
}

func fib_dp_opt(n int, stats *sort.GoStats) uint64  {
	if n <= 1 {
		return 1
	}

	var zero uint64 = 1
	var first uint64 = 1
	var result uint64 = first + zero
	for i := 2; i <= n ; i++ {
		result = first + zero
		zero = first
		first = result
	}
	return result
}

func main() {
	const fib_num = 40
	{
		fmt.Println("递归计算：斐波那契项前(", fib_num, ")值================")
		stats := sort.GoStats{}
		stats.StartStats()
		result := fib_vol(fib_num, &stats)
		stats.StopStats()
		stats.PrintResult()
		fmt.Println("递归计算：斐波那契项前(", fib_num, ")值 : ", result)
	}

	{
		fmt.Println("DP计算：斐波那契项前(", fib_num, ")值================")
		stats := sort.GoStats{}
		stats.StartStats()
		result := fib_dp(fib_num, &stats)
		stats.StopStats()
		stats.PrintResult()
		fmt.Println("DP计算：斐波那契项前(", fib_num, ")值 : ", result)
	}

	{
		fmt.Println("DP计算：斐波那契项前(", fib_num, ")值================")
		stats := sort.GoStats{}
		stats.StartStats()
		result := fib_dp_opt(fib_num, &stats)
		stats.StopStats()
		stats.PrintResult()
		fmt.Println("DP计算：斐波那契项前(", fib_num, ")值 : ", result)
	}
	
}
