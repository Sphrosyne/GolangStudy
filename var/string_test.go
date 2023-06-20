package Test

import (
	"fmt"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	var builder strings.Builder
	builder.WriteString("wz")
	builder.WriteString("-xx")

	fmt.Println(builder.String())

}
