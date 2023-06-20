package Test

import (
	"fmt"
	"testing"
)

type Book struct {
	auth  string
	title string
}

func changBook1(book Book) {
	book.auth = "xxs"

}

func changeBook2(book *Book) {
	book.auth = "ZH"
}

func TestStruct(t *testing.T) {
	var book1 Book
	book1.title = "SxRes"
	book1.auth = "Lem"
	fmt.Println(book1)

	changBook1(book1)
	fmt.Println(book1)

	changeBook2(&book1)
	fmt.Println(book1)

}
