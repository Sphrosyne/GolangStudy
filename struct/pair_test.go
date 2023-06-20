package Test

import (
	"fmt"
	"testing"
)

type Reader interface {
	ReadBook()
}
type Writer interface {
	WriteBook()
}

type Book1 struct {
}

func (this *Book1) ReadBook() {
	fmt.Println("Read a book")
}

func (this *Book1) WriteBook() {
	fmt.Println("Write a book")
}

func TestPair(t *testing.T) {
	book := &Book1{}

	var r Reader
	r = book
	r.ReadBook()

	var w Writer
	w = r.(Writer)
	w.WriteBook()
}
