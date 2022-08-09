package main

import "fmt"

const (
	BEIJING = iota
	SHANGHAI
	SHENZHEN
)

func main() {
	fmt.Println(BEIJING)
	fmt.Println(SHANGHAI)
	fmt.Println(SHENZHEN)

}
