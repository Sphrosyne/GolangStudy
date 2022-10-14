package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
	"yunion.io/x/log"
)

var (
	counter   int32
	waitGroup sync.WaitGroup
)

func add(i int) {
	defer waitGroup.Done()
	time.Sleep(1 * time.Second)
	log.Infof("线程", i, "初始值", counter)
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
	log.Infof("", i, "最终值", counter)
}

func main() {
	runtime.GOMAXPROCS(1)
	waitGroup.Add(2)

	go add(1)
	go add(2)

	waitGroup.Wait()
	fmt.Println(counter)
}
