package ttycap
//What of detecting capability once and caching it

import (
	"os"
	"golang.org/x/term"
)

// hasFlag, isTTY

type Capabilities struct {
	IsTTY  bool
	ColorCap ColorCap
}

type ColorCap int

const (
	ColorNone  ColorCap = iota
	ColorANSI16 
	ColorANSI256 
	ColorANSI16m
)


func QueryColorCap(stream *os.File) ColorCap{
	//return Capabilities{
		//IsTTY: term.IsTerminal(int(stream.Fd())),
		return detectColorCap(stream)
	//}
}


// will focus on capability detection 
// QueryUnicodeCap(*os.File)
//QueryAll(*os.File) //for Capabilities

//IsTYY might get dropped 
func IsTTY(stream *os.File) bool {
	return term.IsTerminal(int(stream.Fd()))
}
