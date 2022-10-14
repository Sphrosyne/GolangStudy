package main

import "fmt"

var bol = true
var a = 3

func main() {
	change(&bol, &a)

	fmt.Println(bol)
	fmt.Println(a)

}
func change(xx *bool, yy *int) {
	*xx = false
	*yy = 4
}
