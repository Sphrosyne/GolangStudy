package main

import (
	"fmt"
	"time"
)

func main() {
	num := make(chan int, 3)
	go func() {
		defer fmt.Println("goroutine 结束")
		fmt.Println("goroutine start....")
		for i := 0; i < 4; i++ {
			num <- i
			fmt.Println("go----", i)
		}
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("main start....")
	for i := 0; i < 4; i++ {
		k := <-num
		fmt.Println("main-----", k)
	}
	//可能主线程已经结束，协程还没执行完毕
	time.Sleep(2 * time.Second)

}
