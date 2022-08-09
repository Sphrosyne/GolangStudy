package main

import (
	"fmt"
	"reflect"
)

func reflectGet(arg interface{}) {
	fmt.Println(reflect.TypeOf(arg))
	fmt.Println(reflect.ValueOf(arg))
}

func main() {
	var a int
	a = 10
	reflectGet(a)
}
