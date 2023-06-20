package Test

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	var myArray [10]int
	for i := 0; i < len(myArray); i++ {
		myArray[i] = 1
	}

	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}

	for index, value := range myArray {
		fmt.Println(index, "--", value)
	}

	for _, value := range myArray {
		fmt.Println(value)
	}
}
