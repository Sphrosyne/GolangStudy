package Test

import (
	"fmt"
	"testing"
)

func TestMap1(t *testing.T) {
	map1 := make(map[string]string, 5)
	map1["Beijing"] = "China"
	map1["DC"] = "USA"
	map1["Tokyo"] = "Japan"

	delete(map1, "DC")

	for key, value := range map1 {
		fmt.Println(key, value)
	}

	var xa []string
	xa = nil
	fmt.Println(xa == nil)
	for i, s := range xa {
		fmt.Println("---", xa[i])
		fmt.Println("---", s)
	}
}
