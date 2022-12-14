package d7cc

import (
	"fmt"
	"strings"
)

// Reset sequence
const RESET = "\u001b[0m"

// Foreground color sequence
const FG_RED = "\033[31m"
const FG_GREEN = "\033[32m"
const FG_YELLOW = "\033[33m"
const FG_BLUE = "\033[34m"
const FG_PURPLE = "\033[35m"
const FG_CYAN = "\033[36m"
const FG_GRAY = "\033[37m"
const FG_WHITE = "\033[97m"

// _B_ means Bright
const FG_B_BLACK = "\u001b[30;1m"
const FG_B_RED = "\u001b[31;1m"
const FG_B_GREEN = "\u001b[32;1m"
const FG_B_YELLOW = "\u001b[33;1m"
const FG_B_BLUE = " \u001b[34;1m"
const FG_B_MAGENTA = "\u001b[35;1m"
const FG_B_CYAN = "\u001b[36;1m"
const FG_B_WHITE = "\u001b[37;1m"

// Background color sequence
const BG_BLACK = "\u001b[40m"
const BG_RED = "\u001b[41m"
const BG_GREEN = "\u001b[42m"
const BG_YELLOW = "\u001b[43m"
const BG_BLUE = "\u001b[44m"
const BG_MAGENTa = "\u001b[45m"
const BG_CYAN = "\u001b[46m"
const BG_WHITE = "\u001b[47m"

// _B_ means Bright
const BG_B_RED = "\u001b[41;1m"
const BG_B_BLACK = "\u001b[40;1m"
const BG_B_GREEN = "\u001b[42;1m"
const BG_B_YELLOW = "\u001b[43;1m"
const BG_B_BLUE = "\u001b[44;1m"
const BG_B_MAGENTa = "\u001b[45;1m"
const BG_B_CYAN = "\u001b[46;1m"
const BG_B_WHITE = "\u001b[47;1m"

// Sequence for text
const TXT_BOLD = "\u001b[1m"
const TXT_UNDERLINE = "\u001b[4m"
const TEX_REVERSED = "\u001b[7m" // change places of FG and BG colors

// Example of color sequence using
// In other project don`t foreget add <package_import_name> (d7cc - default) before sequence code
func ExampleForeGround() {
	builder := strings.Builder{}
	builder.WriteString("This is ")
	builder.WriteString(FG_CYAN)
	builder.WriteString("CYAN")
	builder.WriteString(RESET)
	builder.WriteString(". ")

	builder.WriteString("This is ")
	builder.WriteString(FG_RED)
	builder.WriteString("RED ")
	builder.WriteString(FG_GREEN)
	builder.WriteString(TXT_UNDERLINE)
	builder.WriteString("GREEN")
	builder.WriteString(RESET)

	fmt.Println(builder.String())
}
