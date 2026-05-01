package termcolor

import (
	"os"
)

type ColorCap int

const (
	unknown   ColorCap = -1
	ColorNone  ColorCap = 0
	Color16 ColorCap = 16
	Color256 ColorCap = 256
	ColorTrue ColorCap = 16777216
)




func Capability(stream ...*os.File) (ColorCap, DetectionStatus) {
	targetStream := os.Stdout
	cachedArgs := os.Args[1:]
	if len(stream) > 0 && stream[0] != nil {
		targetStream = stream[0]
	}

	if checkColorToggle(cachedArgs, "color", "=", colorPosArgs) || checkColorToggle(cachedArgs, "color", "=", colorAutoArgs){
		cap, status := detectColorCap(cachedArgs, targetStream)
		if status == Detected {
			return cap, Detected//true
		}
		//Forced color on
		//Terminal may not support ansi16 (dumb terminal). so return uncertain
		return Color16, Uncertain//true
	}
	
	return detectColorCap(cachedArgs, targetStream)
}