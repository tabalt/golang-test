package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	goroutine_sleep()
	fmt.Println("-------")

	goroutine_channel()
	fmt.Println("-------")

	goroutine_waitgroup()
	fmt.Println("-------")

	goroutine_timeout()
	fmt.Println("-------")
}
func goroutine_sleep() {
	for i := 0; i < 10; i++ {
		go func(num int) {
			fmt.Println(num)
		}(i)
	}

	//仅测试，实际不能用
	time.Sleep(time.Second)
}

func goroutine_waitgroup() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int) {
			fmt.Println(num)
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func goroutine_channel() {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go func(num int) {
			c <- num
		}(i)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}

func goroutine_timeout() {
	c := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go func(num int) {
			time.Sleep(time.Duration(num) * time.Second)
			c <- num
		}(i)
	}

	timer := time.NewTimer(time.Duration(6) * time.Second)
loop:
	for i := 0; i < 10; i++ {
		select {
		case num := <-c:
			fmt.Println(num)
		case <-timer.C:
			fmt.Println("timeout break")
			break loop
		}
	}
}

func goroutine_context() {

}
