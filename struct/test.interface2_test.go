package Test

import (
	"fmt"
	"testing"
)

func getArgType(arg interface{}) {
	fmt.Println("入参是", arg)
	value, ok := arg.(string)
	if ok {
		fmt.Println("string 类型", value)
	} else {
		fmt.Println("非string类型", value)
	}
}

func TestInterface2(t *testing.T) {
	var param1 int = 1
	var param2 string = "新车"
	var param3 float32 = 1.32

	getArgType(param1)
	getArgType(param2)
	getArgType(param3)
}
