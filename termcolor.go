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

var (
	colorOff = checkColorToggle("color", "=", colorNegArgs) || checkColorToggle("no-color", "=", colorNegArgs) || checkColorToggle("no-color", "=", colorPosArgs)
	colorOn = checkColorToggle("color", "=", colorPosArgs)
	colorAuto = checkColorToggle("color", "=", colorAutoArgs)
)

//Terminal supports color?
func SupportsColor() bool {
	return (!Disabled() && IsTerminal()) || Enabled() || colorAuto /*|| detectColorCap() > ColorNone */
	//detectColorCap() conflicts with the environment variables as terminal can support colors but NO_COLOR is set
}


func Disabled() bool {
	_, exists := os.LookupEnv("NO_COLOR")
	return exists || colorOff
}

func Enabled() bool {
	_, exists := os.LookupEnv("FORCE_COLOR")
	return exists || colorOn
}

func Capability() ColorCap {
	if !IsTerminal() || Disabled() {
		return ColorNone
	}
	return detectColorCap()
}


func  IsTerminal() bool {
	return term.IsTerminal(int(os.Stdout.Fd()))
}
