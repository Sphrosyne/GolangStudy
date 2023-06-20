package Test

import (
	"fmt"
	"testing"
	"yunion.io/x/jsonutils"
)

func TestSlice2(t *testing.T) {
	//追加、截取、复制
	myArray := make([]int, 3, 5)
	myArray = []int{1, 3, 5}

	myArray = append(myArray, 2)
	//myArray = append(myArray, 7)
	//myArray = append(myArray, 9)
	fmt.Println(len(myArray), cap(myArray))
	myArray2 := myArray[0:2]

	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}
	fmt.Println("============")
	myArray[0] = 9999
	for i := 0; i < len(myArray2); i++ {
		fmt.Println(myArray2[i])
	}

	fmt.Println("===============")
	myArray3 := make([]int, 1, 5)
	copy(myArray3, myArray)
	for i := 0; i < len(myArray3); i++ {
		fmt.Println(myArray3[i])
	}
	fmt.Println("===============")
	var projectIds []string
	ax := []string{"x1", "23s", "asdw"}
	for _, s := range ax {
		projectIds = append(projectIds, s)
	}
	fmt.Println(projectIds)
	fmt.Println(jsonutils.NewStringArray(projectIds))

}
