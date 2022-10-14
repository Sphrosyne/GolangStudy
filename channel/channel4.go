package main

import (
	"fmt"
)

func main() {
	num := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i, "---------", <-num)
		}
		quit <- 0
	}()

	fil(num, quit)

}
func fil(num, quit chan int) {
	x, y := 1, 1
	k := 0
	for {
		k++
		fmt.Println("第", k, "次")
		select {
		case num <- x:
			//如果num可以写，则可以进来
			x = y
			y = x + y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
