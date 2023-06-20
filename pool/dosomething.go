package pool

import (
	"fmt"
	"time"
)

type Dosomething struct {
	Num int
}

func (d *Dosomething) Do() {
	fmt.Println("开启线程：", d.Num)
	time.Sleep(5 * time.Second)
	fmt.Println("线程结束", d.Num)
}
