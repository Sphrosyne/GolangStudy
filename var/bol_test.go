package Test

import (
	"fmt"
	"testing"
)

var bol = true
var a = 3

func TestBol(t *testing.T) {
	change(&bol, &a)

	fmt.Println(bol)
	fmt.Println(a)

}
func change(xx *bool, yy *int) {
	*xx = false
	*yy = 4
}
