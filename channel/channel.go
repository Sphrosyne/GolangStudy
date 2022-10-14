package main

import (
	"fmt"
	"time"
)

func main() {

	num := make(chan int)

	go func() {
		defer fmt.Println("goroutine end ....")
		fmt.Println("goroutine start .....")
		num <- 666
		time.Sleep(4 * time.Second)
	}()

	xx := <-num
	fmt.Println("值", xx)

	//主线程等支线程结束
	time.Sleep(5 * time.Second)

}
