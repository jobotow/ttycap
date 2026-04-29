package termcolor

import (
	"os"
	"golang.org/x/term"
)

type ColorCap int

const (
	unknown   ColorCap = -1
	ColorNone  ColorCap = 0
	Color8 ColorCap = 8
	Color16 ColorCap = 16
	Color256 ColorCap = 256
	ColorTrue ColorCap = 16777216
)

//Terminal supports color?
func SupportsColor() bool {
	return (!IsDisabled() && IsTerminal()) || detectColorCap() > ColorNone
}


func IsDisabled() bool {
	_, exists := os.LookupEnv("NO_COLOR")
	return exists
}

func Capability() ColorCap {
	if !IsTerminal() || IsDisabled() {
		return ColorNone
	}
	return detectColorCap()
}


func  IsTerminal() bool {
	return term.IsTerminal(int(os.Stdout.Fd()))
}