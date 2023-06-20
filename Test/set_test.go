package Test

import (
	"fmt"
	"gopkg.in/fatih/set.v0"
	"testing"
)

func TestSet(t *testing.T) {
	userIdSet := set.New(set.ThreadSafe)
	userIdSet.Add("1123123123123")
	userIdSet.Add("xxxq2e33")
	userIdSet.Add("kiiiahyhasd")
	list := userIdSet.List()
	fmt.Println(userIdSet.Size())
	fmt.Println(list)
	for _, v := range list {
		fmt.Println(v)
	}
}
