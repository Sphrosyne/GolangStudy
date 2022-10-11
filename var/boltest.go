package main

import "fmt"

var bol = true

func main() {

	fmt.Println(bol)
	change(bol)
	fmt.Println(bol)

}
func change(bol bool) {
	bol = false
}
