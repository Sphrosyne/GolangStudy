package main

import (
	"fmt"
	"reflect"
)

type man struct {
	//go语言里面struct里面变量如果大写则是public,如果是小写则是private的，private的时候通过反射不能获取其值
	Name string
	Age  int
	Addr string
}

func (this *man) sleep() {
	fmt.Println("sleep .....")
}

func main() {
	zh := man{"xs", 18, "Beijing"}

	fmt.Println(reflect.TypeOf(zh))

	fmt.Println(reflect.ValueOf(zh))

	rt := reflect.TypeOf(zh)
	rv := reflect.ValueOf(zh)

	for i := 0; i < rt.NumField(); i++ {
		fmt.Println(rt.Field(i).Name, rt.Field(i).Type)
		value := rv.Field(i).Interface()
		fmt.Println(value)

	}
}
