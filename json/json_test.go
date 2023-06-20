package Test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Movie struct {
	Year  string   `json:"year"`
	Title string   `json:"title"`
	Actor []string `json:"actor"`
}

func TestJson(t *testing.T) {
	m := Movie{"2000", "喜剧之王", []string{"周星驰", "张柏芝"}}

	//转json
	res, err := json.Marshal(m)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}
	fmt.Printf("%s\n", res)
	//fmt.Println(" 转json结果", res)

	myMovie := Movie{}
	err = json.Unmarshal(res, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal error", err)
		return
	}
	fmt.Printf("%v\n", myMovie)
}
