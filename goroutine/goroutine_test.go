package Test

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	go func() {
		k := 0
		for {
			k++
			fmt.Println("goroutine2.....", k)
			time.Sleep(1 * time.Second)
		}
	}()

	go task()

	i := 0
	for {
		i++
		fmt.Println("main....", i)
		time.Sleep(1 * time.Second)
	}
}

func task() {
	k := 0
	for {
		k++
		fmt.Println("goroutine.....", k)
		time.Sleep(1 * time.Second)
	}
}
