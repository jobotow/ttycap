# TermColor
**termcolor** tells you what colors your terminal can display and to what level or depth

**It answers questions like**
- Does my terminal supports colors at all?
- Can it show 256 colors?
- Can it show truecolor

Use this library to decide what color codes to send to the terminal


## What is does not do
- It does NOT print colored text
- It does NOT format strings
- It only detects terminal color capability.


## Installation
```bash
go get -u github.com/phamio/termcolor
```


# Example1
```go
package main
import (
    "fmt"
    "github.com/phamio/termcolor"
)

func main(){
    caps := termcolor.Capability()

    if caps == termcolor.ColorTrue {
        fmt.Println("Terminal supports true color")
    } else if caps == termcolor.Color256 {
        fmt.Println("Terminal supports only 256-color palette")
    } else if caps == termcolor.Color16 {
        fmt.Println("Terminal supports only 16 colors")
    } else if caps == termcolor.Color8 {
        fmt.Println("Terminal supports only 8 colors")
    } else if caps == termcolor.ColorNone {
        fmt.Println("Terminal supports no color (dumb terminal)")
    } else {
        fmt.Println("Done")
    }
}
```

# Example2

```go
package main
import (
    "fmt"
    "github.com/phamio/termcolor"
)

func main(){
    //Check if output is to terminal
    term := termcolor.IsTerminal()
    if term {
        fmt.Println("Output is to terminal")
    }

    //Check if terminal color is disabled
    termCap := termcolor.IsDisabled()
    if termCap {
        fmt.Prinln("Terminal color is disabled")
    }

    //Check if terminal supports color
    termCol := termcolor.SupportsColor()
    if termCol {
        fmt.Println("Hurray! Terminal understand colors")
    }
}
