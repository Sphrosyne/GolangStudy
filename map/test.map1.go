package main

import "fmt"

func main() {
	map1 := make(map[string]string, 5)
	map1["Beijing"] = "China"
	map1["DC"] = "USA"
	map1["Tokyo"] = "Japan"

	delete(map1, "DC")

	for key, value := range map1 {
		fmt.Println(key, value)
	}
}
