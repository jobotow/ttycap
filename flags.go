package termcolor

import (
	"strings"
	"os"
)

var (
	colorPosArgs = []string{"on", "true", "yes", "1", "always", "force"}
	colorNegArgs = []string{"off", "false", "no", "0", "never", "none"}
	colorAutoArgs = []string{"auto", "tty", "detect"}
)

func matchNextArg(arg string, seqArgs []string) bool {
	for _, value := range seqArgs {
		if arg == value {
			return true
		}
	}
	return false
}


func checkColorToggle(flag, delimiter string, seqArgs []string) bool {
	nextArg, boolean := getNextArg(flag, delimiter)
	if boolean{
		return matchNextArg(nextArg, seqArgs)
	}
	return false
}


func getNextArg(flag, delimiter string) (string, bool) {
	flagDelim := flag+delimiter
	args := os.Args[1:]
	for i, arg := range args{
		cleanArg := strings.TrimLeft(arg, "-")
		//GNU Style
		if strings.HasPrefix(cleanArg, flagDelim){
			return strings.TrimPrefix(cleanArg, flagDelim), true
		}

        //POSIX Style
		if cleanArg == flag {
		    if i+1 < len(args) {
				nextArg := args[i+1]
				//Check if next arg is not a flag
				if !strings.HasPrefix(nextArg, "-"){
					//Thats the value of given flag
					return nextArg, true
				}
			}
			//if flag has no value (bool)
			return "", true
		}
	}
	//Flag wasn't present
	return "", false
}