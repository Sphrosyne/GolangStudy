package main

import "fmt"

func foo1(a int, b int) int {
	fmt.Println(a)
	fmt.Println(b)
	return 2000
}

func foo2(a int, b string) (string, int) {
	fmt.Println(a)
	fmt.Println(b)

	return "xs", 221
}

func foo3(a int, b string) (r1 string, r2 float32) {
	fmt.Println(a)
	fmt.Println(b)

	r1 = "ouu"
	r2 = 1.2
	return
}

func main() {
	var c = foo1(100, 102)
	fmt.Println(c)

	var res1, res2 = foo2(101, "wwd")
	fmt.Println(res1)
	fmt.Println(res2)

	var r1, r2 = foo3(333, "oui")

	fmt.Println(r1)
	fmt.Println(r2)

}
