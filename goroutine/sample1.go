package main

import (
	"fmt"
	"time"
)

func testfrom(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, " : ", i)
	}
}

func testchannel(msgchan <-chan string) {
	msg := <-msgchan
	fmt.Println("recv from other : ", msg)
}

func work(done chan<- bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func ping(ping chan<- string, msg string) {
	ping <- msg

}

func pong(ping <-chan string, pong chan<- string) {
	msg := <-ping
	pong <- msg
}

func testtimeout(c chan<- string, result string, timeoutms int) {
	time.Sleep(time.Duration(timeoutms) * time.Millisecond)
	c <- result
}

func main() {
	{
		c1 := make(chan string)
		go testtimeout(c1, "result one", 500)

		select {
		case msg := <-c1:
			fmt.Println(msg)
		case <-time.After(time.Duration(100) * time.Millisecond):
			fmt.Println("time out")
		default:
			fmt.Println("no task")
		}
		return
	}

	{
		c1 := make(chan string)
		c2 := make(chan string)

		go func(c chan<- string) {
			time.Sleep(time.Second)
			c <- "from one"
		}(c1)

		go func(c chan<- string) {
			time.Sleep(time.Second * 2)
			c <- "from two"
		}(c2)

		for i := 0; i < 2; i++ {
			select {
			case msg := <-c1:
				fmt.Println(msg)
			case msg := <-c2:
				fmt.Println(msg)
			case <-time.After(time.Millisecond * 500):
				fmt.Println("time out")
			}
		}

	}

	//{
	//	ci := make(chan int)
	//
	//	go func() {
	//		fmt.Println(<-ci)
	//	}()
	//	ci <- 88
	//	time.Sleep(time.Second)
	//}

	//{
	//	pings := make(chan string)
	//	pongs := make(chan string)
	//	go ping(pings, "passes message")
	//	go pong(pings, pongs)
	//	fmt.Println(<-pongs)
	//}

	//{
	//	testfrom("direct")
	//
	//	go testfrom("goroutine")
	//
	//	go func(msg string) {
	//		fmt.Println(msg)
	//	}("going")
	//}
	//
	//{
	//	fmt.Println("test channel")
	//	strchanl := make(chan string)
	//
	//	go testchannel(strchanl)
	//	var str string
	//	fmt.Scanln(&str)
	//	strchanl <- str
	//
	//	time.Sleep(time.Second)
	//
	//}

	//{
	//	message := make(chan string, 2)
	//	message <- "alexi"
	//	message <- "chen"
	//
	//	fmt.Println(<-message)
	//	fmt.Println(<-message)
	//
	//}
	//
	//{
	//	done := make(chan bool)
	//	go work(done)
	//
	//	<-done
	//}

}
