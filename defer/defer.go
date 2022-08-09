package main

import "fmt"

func returnFunc() int {
	fmt.Println("returnFunc call")
	return 0
}
func deferFunc() int {
	fmt.Println("deferFunc call")
	return 0
}
func test() int {
	defer deferFunc()
	return returnFunc()
}

func main() {

	defer fmt.Println("b")
	defer fmt.Println("C")
	fmt.Println("A")
	
	test()
}
