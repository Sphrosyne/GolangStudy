package Test

import (
	"fmt"
	"testing"
)

type Animal interface {
	sleep()
	getColor() string
	getType() string
}

type cat struct {
	color string
}

func (this *cat) sleep() {
	fmt.Println("cat sleep...")
}
func (this *cat) getColor() string {
	fmt.Println("猫的颜色")
	return "灰色"
}

func (this *cat) getType() string {
	fmt.Println("猫的类型")
	return "卡机"
}

func TestInterface(t *testing.T) {
	var animal Animal
	animal = &cat{
		"红色",
	}
	fmt.Println(animal.getColor())

	animal.sleep()
}
