package main

import "fmt"

func main() {
	/*动态数组的实现方式*/
	//第一种
	var myArray []int
	myArray = make([]int, 3)

	//第二种
	myArray2 := []int{1, 3, 5}

	//第三种类
	myArray3 := make([]int, 5)

	//第四种
	var myArray4 []int = make([]int, 4)

	if myArray4 == nil {
		fmt.Println("xxxx")
		printArray(myArray4)
	} else {
		fmt.Println("yyyy")
	}

	printArray(myArray)
	printArray(myArray2)
	fmt.Println(myArray3)
}

func printArray(myArray []int) {
	for index, value := range myArray {
		fmt.Println(index, value)
	}
}
