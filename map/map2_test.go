package Test

import (
	"fmt"
	"testing"
)

func TestMap3(t *testing.T) {
	projectLeaderMap := make(map[string]map[string]string, 16)
	m := projectLeaderMap["1"]
	if m == nil {
		m = make(map[string]string)
	}
	m["1"] = "x"
	projectLeaderMap["1"] = m
	/*m := make(map[string]string)
	m["1"] = "x"
	m["2"] = "y"
	*/

	m2 := projectLeaderMap["2"]
	if m2 == nil {
		m2 = make(map[string]string)
	}
	m2["2"] = "z"
	projectLeaderMap["2"] = m2
	fmt.Println(projectLeaderMap["1"]["1"])
	fmt.Println(projectLeaderMap["2"]["2"])

	projectLeaderMap["1"]["2"] = "asdwee"
	fmt.Println(projectLeaderMap["1"]["2"])

}

func TestMap4(t *testing.T) {
	projectLeaderMap := make(map[string]map[string]string, 16)
	m := projectLeaderMap["1"]
	if m == nil {
		m = make(map[string]string)
		projectLeaderMap["1"] = m
	}
	projectLeaderMap["1"]["1"] = "x"
	projectLeaderMap["1"]["2"] = "y"
	fmt.Println(projectLeaderMap["1"]["2"])
}
