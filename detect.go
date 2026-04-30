package termcolor

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

	var colorTerm = []string{"screen", "xterm", "vt100", "color", "ansi", "cygwin", "linux"}

	



func detectColorCap() ColorCap {
	if cap, ok := checkTermEnv(); ok {
		return cap
	}

	if cap, ok := queryTput(); ok {
		return cap
	}
	
	//Failed to detect
	return ColorNone
}

func checkTermEnv() (ColorCap, bool) {
	term := os.Getenv("TERM")
	colorTerm := os.Getenv("COLORTERM")

	//Check for explicit color number in $TERM or $COLORTERM
	if supportsTrueColor(colorTerm) {
		return ColorTrue, true
	}
	if supports256Color(term){
		return Color256, true
	}
	if supports16Color(term){
		return Color16, true
	}

	if supportsNone(term){
		return ColorNone, true
	}

	//Detection failed
	return unknown, false
}

      



func queryTput() (ColorCap, bool) {
	//Check if tput exists
	_, err := exec.LookPath("tput")
	if err != nil {
		//Detection failed
		return unknown, false
	}
    
	//Run tput colors
	cmd := exec.Command("tput", "colors")
	out, err := cmd.Output()
	if err != nil {
		return unknown, false
	}
	
	colors, err := strconv.Atoi(strings.TrimSpace(string(out)))
	if err != nil || colors <= 0 {
		return unknown, false
	}

	return ColorCap(colors), true
}







func supportsTrueColor(colorTerm string) bool {
	return colorTerm == "truecolor" || colorTerm == "24bit"
}

func supports256Color(term string) bool {
	return strings.Contains(term, "256color")
}

func supports16Color(term string) bool {
	return strings.Contains(term, "16color") || supportsColor(term)
}

func supportsColor(term string) bool {
	for _, t := range colorTerm {
		if strings.Contains(term, t) {
			return true
		}
	}
	return false
}
 
func supportsNone(term string) bool {
	return term == "dumb" || strings.Contains(term, "mono")
}


//WIP
//func windowsColorSupport() bool {
//	_, exists := os.LookupEnv("ANSICON")
//	return exists
//}


