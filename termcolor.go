package termcolor
//What of detecting capability once and caching it

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

type Option func(*config)
type  config struct {
	stream *os.File
	sniffFlags bool
}

func defaultConfig() config {
	return config{
		stream: os.Stdout,
		sniffFlags: true,
	}
}

func Stream(s *os.File) Option{
	return func(c *config) {
		c.stream = s 
	}
}

func SniffFlagToggle(sniff bool) Option {
	return func(c *config) {
		c.sniffFlags = sniff
	}
}


func Capability(opts ...Option) (ColorCap, bool){
	cachedArgs := os.Args[1:]
	cfg := defaultConfig()
	for _, opt := range opts {
		opt(&cfg)
	}
	if cfg.sniffFlags {
		getColorFlagCap(cachedArgs, cfg)
	}
	return detectColorCap(cachedArgs, cfg)
}





func getColorFlagCap(cachedArgs []string, cfg config)(ColorCap, bool) {
	if checkColorToggle(cachedArgs, "color", "=", colorPosArgs) || checkColorToggle(cachedArgs, "color", "=", colorAutoArgs){
		cap, ok := detectColorCap(cachedArgs, cfg)
		if ok {
			return cap, true
		}
		//Forced color on
		return Color16, true
	} 
	
	//Check --no-color flags
    if checkColorToggle(cachedArgs, "no-color", "=", colorNegArgs) || checkColorToggle(cachedArgs, "color", "=", colorNegArgs) || checkColorToggle(cachedArgs,"no-color", "=", colorPosArgs) {
	    return ColorNone, true
	}
	return ColorNone, true

}