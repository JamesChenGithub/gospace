package sort

import (
	"fmt"
	"time"
)

type GoStats struct  {
	CompareCount int
	SwapCount int
	StartTime time.Time
	Elapsed time.Duration
}

func (stats *GoStats)StartStats()  {
	stats.StartTime = time.Now()
}

func (stats *GoStats)StopStats(){
	stats.Elapsed = time.Since(stats.StartTime)
}

func (stats GoStats)PrintResult() {
	fmt.Println("共比较 : ", stats.CompareCount, " 次")
	fmt.Println("共交互 : ", stats.SwapCount, " 次")
	fmt.Println("总耗时 : ", stats.Elapsed, " 次")
}