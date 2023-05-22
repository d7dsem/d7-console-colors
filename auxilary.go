package d7cc

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/sys/windows"
)

// Reset sequence
var RESET = "\u001b[0m"

// Foreground color sequence
var FG_RED = "\033[31m"
var FG_GREEN = "\033[32m"
var FG_YELLOW = "\033[33m"
var FG_BLUE = "\033[34m"
var FG_PURPLE = "\033[35m"
var FG_CYAN = "\033[36m"
var FG_GRAY = "\033[37m"
var FG_WHITE = "\033[97m"

// _B_ means Bright
var FG_B_BLACK = "\u001b[30;1m"
var FG_B_RED = "\u001b[31;1m"
var FG_B_GREEN = "\u001b[32;1m"
var FG_B_YELLOW = "\u001b[33;1m"
var FG_B_BLUE = " \u001b[34;1m"
var FG_B_MAGENTA = "\u001b[35;1m"
var FG_B_CYAN = "\u001b[36;1m"
var FG_B_WHITE = "\u001b[37;1m"

// Background color sequence
var BG_BLACK = "\u001b[40m"
var BG_RED = "\u001b[41m"
var BG_GREEN = "\u001b[42m"
var BG_YELLOW = "\u001b[43m"
var BG_BLUE = "\u001b[44m"
var BG_MAGENTa = "\u001b[45m"
var BG_CYAN = "\u001b[46m"
var BG_WHITE = "\u001b[47m"

// _B_ means Bright
var BG_B_RED = "\u001b[41;1m"
var BG_B_BLACK = "\u001b[40;1m"
var BG_B_GREEN = "\u001b[42;1m"
var BG_B_YELLOW = "\u001b[43;1m"
var BG_B_BLUE = "\u001b[44;1m"
var BG_B_MAGENTa = "\u001b[45;1m"
var BG_B_CYAN = "\u001b[46;1m"
var BG_B_WHITE = "\u001b[47;1m"

// Sequence for text
var TXT_BOLD = "\u001b[1m"
var TXT_UNDERLINE = "\u001b[4m"
var TEX_REVERSED = "\u001b[7m" // change places of FG and BG colors

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

func Colorize(d interface{}, color string) string {
	s := strings.Builder{}
	s.WriteString(color)
	s.WriteString(fmt.Sprintf("%v", d))
	s.WriteString(RESET)
	return s.String()
}

func TurnOnColorizing() {
	// Reset sequence
	RESET = "\u001b[0m"

	// Foreground color sequence
	FG_RED = "\033[31m"
	FG_GREEN = "\033[32m"
	FG_YELLOW = "\033[33m"
	FG_BLUE = "\033[34m"
	FG_PURPLE = "\033[35m"
	FG_CYAN = "\033[36m"
	FG_GRAY = "\033[37m"
	FG_WHITE = "\033[97m"

	// _B_ means Bright
	FG_B_BLACK = "\u001b[30;1m"
	FG_B_RED = "\u001b[31;1m"
	FG_B_GREEN = "\u001b[32;1m"
	FG_B_YELLOW = "\u001b[33;1m"
	FG_B_BLUE = " \u001b[34;1m"
	FG_B_MAGENTA = "\u001b[35;1m"
	FG_B_CYAN = "\u001b[36;1m"
	FG_B_WHITE = "\u001b[37;1m"

	// Background color sequence
	BG_BLACK = "\u001b[40m"
	BG_RED = "\u001b[41m"
	BG_GREEN = "\u001b[42m"
	BG_YELLOW = "\u001b[43m"
	BG_BLUE = "\u001b[44m"
	BG_MAGENTa = "\u001b[45m"
	BG_CYAN = "\u001b[46m"
	BG_WHITE = "\u001b[47m"

	// _B_ means Bright
	BG_B_RED = "\u001b[41;1m"
	BG_B_BLACK = "\u001b[40;1m"
	BG_B_GREEN = "\u001b[42;1m"
	BG_B_YELLOW = "\u001b[43;1m"
	BG_B_BLUE = "\u001b[44;1m"
	BG_B_MAGENTa = "\u001b[45;1m"
	BG_B_CYAN = "\u001b[46;1m"
	BG_B_WHITE = "\u001b[47;1m"

	// Sequence for text
	TXT_BOLD = "\u001b[1m"
	TXT_UNDERLINE = "\u001b[4m"
	TEX_REVERSED = "\u001b[7m" // change places of FG and BG colors
}

func TurnOffColorizing() {
	// Reset sequence
	RESET = ""

	// Foreground color sequence
	FG_RED = ""
	FG_GREEN = ""
	FG_YELLOW = ""
	FG_BLUE = ""
	FG_PURPLE = ""
	FG_CYAN = ""
	FG_GRAY = ""
	FG_WHITE = ""

	// _B_ means Bright
	FG_B_BLACK = ""
	FG_B_RED = ""
	FG_B_GREEN = ""
	FG_B_YELLOW = ""
	FG_B_BLUE = ""
	FG_B_MAGENTA = ""
	FG_B_CYAN = ""
	FG_B_WHITE = ""

	// Background color sequence
	BG_BLACK = ""
	BG_RED = ""
	BG_GREEN = ""
	BG_YELLOW = ""
	BG_BLUE = ""
	BG_MAGENTa = ""
	BG_CYAN = ""
	BG_WHITE = ""

	// _B_ means Bright
	BG_B_RED = ""
	BG_B_BLACK = ""
	BG_B_GREEN = ""
	BG_B_YELLOW = ""
	BG_B_BLUE = ""
	BG_B_MAGENTa = ""
	BG_B_CYAN = ""
	BG_B_WHITE = ""

	// Sequence for text
	TXT_BOLD = ""
	TXT_UNDERLINE = ""
	TEX_REVERSED = "" // change places of FG and BG colors
}

func TurnOnVirtualTerminalProcesingWindows() {
	stdout := windows.Handle(os.Stdout.Fd())
	var originalMode uint32

	windows.GetConsoleMode(stdout, &originalMode)
	windows.SetConsoleMode(stdout, originalMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
}

func FormForegroundRGBColor(r, g, b int) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

func FormBackroundRGBColor(r, g, b int) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}
