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


func Query(stream *os.File) Capabilities{
	return Capabilities{
		IsTTY: term.IsTerminal(int(stream.Fd())),
		ColorCap: detectColorCap(stream)
	}
}


func IsTTY(stream *os.File) bool {
	return term.IsTerminal(int(stream.Fd()))
}
