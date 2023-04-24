package Test

import (
	"fmt"
	"gopkg.in/fatih/set.v0"
	"strings"
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

func TestSet2(t *testing.T) {
	s := "axasasd"
	split := strings.Split(s, ",")
	fmt.Println(len(split))

	strins := make([]string, 0)
	for i, strin := range split {
		fmt.Println("******", split[i])
		strins = append(strins, strin)
	}
	fmt.Println("&&&&", len(strins))
	fmt.Println(strins)
}
