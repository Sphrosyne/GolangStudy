package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//readFile()
	eval(1, 2, "+")
}

//读文件
func readFile() {
	const filename = "D:\\文件下载\\setup.log"
	// Go 函数可以返回两个值
	// func ReadFile(filename string) ([]byte, error)
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		// contents 是 []byte, 用%s 可以打印出来
		fmt.Printf("%s", contents)
	}
	// if 语句外部可访问
	fmt.Printf("%s", contents)
}

//switch的使用
func eval(a, b int, op string) {
	var value int
	switch op {
	case "+":
		value = a + b
		//穿透下去
		fallthrough
	case "(":
		fmt.Println("xxxxx")
	case "-":
		value = a - b
	default:
		value = 0
	}
	fmt.Println(value)
}
