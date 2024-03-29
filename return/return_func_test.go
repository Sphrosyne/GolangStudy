package Test

import (
	"fmt"
	"testing"
)

func adder() func(x int) int {
	sum := 0 // 自由变量，Go 的闭包是可以含有自由变量的
	return func(value int) int {
		sum += value
		return sum
	}
}

func TestReturnFunc(t *testing.T) {
	add := adder() // 此处返回一个函数
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + ... + %d = %d", i, add(i)) //调用函数add
		fmt.Println()
	}
}
