package main

import "fmt"

func main() {
	/*map的初始化*/

	//第一种
	var map1 map[string]string
	//第二种
	var map2 map[string]string = make(map[string]string, 5)

	//第三种
	map3 := make(map[string]string, 3)

	//第四种
	map4 := map[string]string{
		"one":   "java",
		"two":   "go",
		"three": "python",
	}

	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)
	fmt.Println(map4)
}
