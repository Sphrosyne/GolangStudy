package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (this *Human) eat(name string) {
	fmt.Println(name, " eat....")
}

func (this *Human) walk() {
	fmt.Println("Human walk.....")
}

type Superman struct {
	human Human
	level int
}

func (this *Superman) Fly() {
	fmt.Println("Superman 飞.....")
}

func main() {
	human := Human{"zhangSan", 6}
	fmt.Println(human)

	superman := Superman{
		Human{"lisi", 6},
		1,
	}
	superman.Fly()
	superman.human.eat(superman.human.name)

	fmt.Println(superman.human.age)
}
