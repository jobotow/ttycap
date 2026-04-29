package termcolor

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)



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
	
	if !supportsNone(term){
		return Color8, true
	}

	if supportsNone(term){
		return ColorNone, true
	}

	if windowsColorSupport(){
		return Color8, true
	}
	//Detection failed
	return unknown, false
}



func supportsTrueColor(colorTerm string) bool {
	return colorTerm == "truecolor" || colorTerm == "24bit"
}

func supports256Color(term string) bool {
	return strings.Contains(term, "256color")
}

func supports16Color(term string) bool {
	return strings.Contains(term, "16color")
}

func supportsNone(term string) bool {
	return term == "dumb" || strings.Contains(term, "mono")
}

func windowsColorSupport() bool {
	_, exists := os.LookupEnv("ANSICON")
	return exists
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