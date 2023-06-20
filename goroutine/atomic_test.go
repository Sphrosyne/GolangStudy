package Test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

var (
	counter   int32
	waitGroup sync.WaitGroup
)

func add(i int) {
	defer waitGroup.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("线程", i, "初始值", counter)
	for count := 0; count < 10000000; count++ {
		value := counter
		// 当前的 goroutine 主动让出资源，从线程退出，并放回到队列，
		// 让其他 goroutine 进行执行
		//runtime.Gosched()

		if count == 100 && i == 1 {
			time.Sleep(1 * time.Second)
		}
		value++
		counter = value

		//atomic.AddInt32(&counter, 1)
	}
	fmt.Println("", i, "最终值", counter)
}

func TestAtomic(t *testing.T) {
	runtime.GOMAXPROCS(1)
	waitGroup.Add(2)

	go add(1)
	go add(2)

	waitGroup.Wait()
	fmt.Println(counter)
}
