package _func

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	fmt.Println(add(1, 2))

	//匿名函数
	var add2 = func(a, b int) int {
		return a + b
	}
	fmt.Println(add2(3, 5))

	fmt.Println(swap(6, 8))

	fmt.Println(sum(1, 2, 3, 4, 5))

	//当可变参数是一个空接口类型时，调用者是否解包可变参数会导致不同的结果：
	var sp = []interface{}{123, "abc"}
	fmt.Println(sp)
	fmt.Println(sp...)

	v := inc(1)
	fmt.Println(v)

}

func add(a, b int) int {
	return a + b
}

//函数返回多个值
func swap(a, b int) (int, int) {
	return b, a
}

//可变数量 的 参数
func sum(a int, b ...int) int {
	for _, v := range b {
		a += v
	}
	return a
}

//返回值也可以命名
func inc(i int) (v int) {
	defer func() { v++ }()
	return i + 10
}
