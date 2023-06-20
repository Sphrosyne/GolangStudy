package Test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var (
	flag       int32
	waitGroup2 sync.WaitGroup
)

func shutDown(name string) {
	defer waitGroup2.Done()
	for {
		time.Sleep(1 * time.Second)
		//if flag == 1 {
		if atomic.LoadInt32(&flag) == 1 {
			fmt.Println(name, "关机")
			break
		}
	}
}

func TestAtomic2(t *testing.T) {

	waitGroup2.Add(1)

	go shutDown("A")

	time.Sleep(3 * time.Second)
	//flag = 1
	atomic.StoreInt32(&flag, 1)
	waitGroup2.Wait()
}
