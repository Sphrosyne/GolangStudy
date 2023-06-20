package Test

import (
	"fmt"
	"testing"
)

const (
	BEIJING = iota
	SHANGHAI
	SHENZHEN
)

func TestConst(t *testing.T) {
	fmt.Println(BEIJING)
	fmt.Println(SHANGHAI)
	fmt.Println(SHENZHEN)

}
