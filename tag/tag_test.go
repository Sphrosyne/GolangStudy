package Test

import (
	"fmt"
	"reflect"
	"testing"
)

type people struct {
	Name string `info:"name" doc:"你的名字"`
	Age  int    `info:"年龄"`
}

func TestTag(t *testing.T) {
	p := people{"zhangSan", 17}
	elem := reflect.TypeOf(&p).Elem()
	for i := 0; i < elem.NumField(); i++ {
		fmt.Println(elem.Field(i).Tag.Get("doc"))
		fmt.Println(elem.Field(i).Tag.Get("info"))
	}
}
