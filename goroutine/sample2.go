package main

import (
	"fmt"
	"time"
)

func doJob(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("Worker", id, "process job", j)
		time.Sleep(time.Second)
		result <- j * 2
	}
}

func main() {
	{
		jobs := make(chan int, 100)
		results := make(chan int, 100)

		for w := 1; w < 3; w++ {
			go doJob(w, jobs, results)
		}

		for j := 1; j < 10; j++ {
			jobs <- j
		}
		close(jobs)

		for a := 1; a < 10; a++ {
			fmt.Println("result : " , <-results)
		}

	}
	return
	{
		ticker := time.NewTicker(time.Millisecond * 500)
		go func() {
			for t := range ticker.C {
				fmt.Println("tick at :", t)
			}
		}()

		time.Sleep(time.Second * 3)
		ticker.Stop()

	}
	{
		done := make(chan bool)

		var timer = time.NewTimer(time.Second)
		<-timer.C
		fmt.Println("timer expired")

		timer2 := time.NewTimer(time.Second)
		go func() {
			<-timer2.C
			fmt.Println("timer2  expired")
			done <- true
		}()

		stop := timer2.Stop()
		if stop {
			fmt.Println("timer2 stoped")
			go func() {
				done <- true
			}()
		}

		<-done
	}

	{
		queue := make(chan string, 2)
		queue <- "one"
		queue <- "two"
		close(queue)

		for elem := range queue {
			fmt.Println(elem)
		}

	}
	{

		jobs := make(chan int, 5)
		done := make(chan bool)

		go func() {
			for {
				j, ok := <-jobs
				if ok {
					fmt.Println("recv jod : ", j)
				} else {
					fmt.Println("recv all jobs")
					done <- true
					return
				}
			}
		}()

		for j := 0; j < 3; j++ {
			jobs <- j
		}
		close(jobs)
		fmt.Println("send all jobs")
		<-done

	}
}
