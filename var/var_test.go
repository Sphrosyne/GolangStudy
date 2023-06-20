package Test

import (
	"fmt"
	"testing"
)

//全局变量
var e string = "toDie"

func TestVar(t *testing.T) {
	//第一种
	var a int
	fmt.Println(a)

	//第二种
	var b int = 100
	fmt.Println(b)

	//第三种
	var c = 200
	fmt.Println(c)

	//第四种
	d := 100
	fmt.Printf("d = %d, d type = %T\n", d, d)

	fmt.Println(e)

	var xx, yy = 11, 22
	fmt.Println(xx, yy)

	var mm, nn = 33, "dongsahn"
	fmt.Println(mm, nn)
}
