package main

import (
	"GolangStudy/pool"
	"fmt"
	"runtime"
	"time"
)

func main() {
	/*//设置最大线程数
	num := 3

	// 注册工作池，传入任务
	// 参数1 初始化worker(工人)并发个数
	p := pool.NewWorkerPool(num)
	p.Run() //有任务就去做，没有就阻塞，任务做不过来也阻塞*/
	p := pool.GetPool()

	//datanum := 100 * 100 * 100 * 100    //模拟百万请求
	datanum := 10
	go func() { //这是一个独立的协程 保证可以接受到每个用户的请求
		for i := 1; i <= datanum; i++ {
			sc := &Dosomething{Num: i}
			p.JobQueue <- sc //往线程池 的通道中 写参数   每个参数相当于一个请求  来了100万个请求
		}
	}()

	k := 1
	for { //阻塞主程序结束
		fmt.Println("runtime.NumGoroutine() :", runtime.NumGoroutine())
		time.Sleep(2 * time.Second)
		k++
		/*if k == 10 {
			go func() {
				p := pool.NewWorkerPool(20)
				p.Run() //有任务就去做，没有就阻塞，任务做不过来也阻塞
			}()
		}
		if k == 12 {
			go func() {
				p := pool.NewWorkerPool(20)
				p.Run() //有任务就去做，没有就阻塞，任务做不过来也阻塞
			}()
		}*/
	}
}
