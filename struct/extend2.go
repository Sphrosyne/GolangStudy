package main

import "fmt"

type User struct {
	name string
	age  int
}

func (user *User) getName() string {
	return user.name
}
func (user *User) getAge() int {
	return user.age
}

type Admin struct {
	//嵌入，类似继承
	User
	role string
}

func main() {
	user := User{"zx", 1}
	fmt.Println(user.age)

	admin := Admin{user, "stuff"}
	fmt.Println(admin.getAge())
}
