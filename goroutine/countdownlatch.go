package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
	"yunion.io/x/jsonutils"
)

func main() {
	// 1. 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(1)

	// 2. 设定等待器，类比 Java CountDownLatch
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	fmt.Println("=== start ===")
	// 3. 创建第一个 goroutine
	/*go func() {
		defer waitGroup.Done() // CountDownLatch#countDown()

		for count := 0; count < 3; count++ {
			fmt.Println("X", count)
			time.Sleep(4 * time.Second)
		}
	}()

	// 4. 创建第二个 goroutine
	go func() {
		defer waitGroup.Done() // CountDownLatch#countDown()

		for count := 0; count < 3; count++ {
			fmt.Println("Y", count)
			time.Sleep(1 * time.Second)
		}
	}()*/

	dict := jsonutils.NewDict()
	fmt.Println("---", dict.String())
	go func1(&waitGroup, dict)
	go func2(&waitGroup, dict)

	// 5. 阻塞 main goroutine
	waitGroup.Wait() // CountDownLatch#await()
	fmt.Println("=== end ===")
	fmt.Println("***", dict.String())

}
func func1(waitGroup *sync.WaitGroup, dict *jsonutils.JSONDict) {
	defer waitGroup.Done() // CountDownLatch#countDown()

	for count := 0; count < 3; count++ {
		fmt.Println("X", count)
		time.Sleep(4 * time.Second)
	}
	dict.Set("par1", jsonutils.NewString("ouy"))
}

func func2(waitGroup *sync.WaitGroup, dict *jsonutils.JSONDict) {
	defer waitGroup.Done() // CountDownLatch#countDown()

	for count := 0; count < 3; count++ {
		fmt.Println("Y", count)
		time.Sleep(1 * time.Second)
	}
	dict.Set("par2", jsonutils.NewString("poux"))
}
