package main

import (
	"fmt"
)

func main() {
	num := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			num <- i
		}

		//close(num)
	}()

	/*for {
		data, ok := <-num
		if ok {
			fmt.Println(data)
		} else {
			break
		}
	}*/

	for data := range num {
		fmt.Println(data)
	}

	fmt.Println("main end.....")

}
