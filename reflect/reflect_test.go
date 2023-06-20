package Test

import (
	"fmt"
	"reflect"
	"testing"
)

func reflectGet(arg interface{}) {
	fmt.Println(reflect.TypeOf(arg))
	fmt.Println(reflect.ValueOf(arg))
}

func TestReflect(t *testing.T) {
	var a int
	a = 10
	reflectGet(a)
}
