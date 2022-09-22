package main

import (
	"fmt"
	"strings"
)

func main() {
	var builder strings.Builder
	builder.WriteString("wz")
	builder.WriteString("-xx")

	fmt.Println(builder.String())

}
