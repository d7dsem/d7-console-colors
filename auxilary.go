package d7cc

import (
	"fmt"
	"strings"
)

var FG_Reset = "\033[0m"
var FG_Red = "\033[31m"
var FG_Green = "\033[32m"
var FG_Yellow = "\033[33m"
var FG_Blue = "\033[34m"
var FG_Purple = "\033[35m"
var FG_Cyan = "\033[36m"
var FG_Gray = "\033[37m"
var FG_White = "\033[97m"

func ExampleForeGround() {
	builder := strings.Builder{}
	builder.WriteString("This is ")
	builder.WriteString(FG_Cyan)
	builder.WriteString("CYAN")
	builder.WriteString(FG_Reset)
	builder.WriteString(". ")

	builder.WriteString("This is ")
	builder.WriteString(FG_Red)
	builder.WriteString("RED ")
	builder.WriteString(FG_Green)
	builder.WriteString("GREEN")
	builder.WriteString(FG_Reset)

	fmt.Println(builder.String())
}
