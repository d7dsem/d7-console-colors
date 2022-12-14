package d7consolecolors

import (
	"fmt"
	"strings"
)

func Example() {
	builder := strings.Builder{}
	// builder.WriteString(color.Cyan)
	// builder.WriteString("Account info:\n")
	// builder.WriteString(color.Reset)

	builder.WriteString("This is test string!")

	fmt.Println(builder.String())
}
