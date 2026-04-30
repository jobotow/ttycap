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


func checkColorToggle(flag, delimiter string, seqArgs []string) bool {
	nextArg, boolean := GetNextArg(flag, delimiter)
	if boolean{
		if nextArg != ""{
		return MatchNextArg(nextArg, seqArgs)
		}
		return true
	}
	return false
}


//=================================
// PUBLIC API FOR FLAGS PARSING
//=================================

func MatchNextArg(arg string, seqArgs []string) bool {
	for _, value := range seqArgs {
		if arg == value {
			return true
		}
	}
	return false
}



func GetNextArg(flag string, delimiter ...string) (string, bool) {
	delim := "="
	if len(delimiter) > 0 {
		delim = delimiter[0]
	}
	flagDelim := flag+delim
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