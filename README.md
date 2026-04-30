# TermColor
**termcolor** tells you what colors your terminal can display and at what level or depth

**It answers questions like**
- Does my terminal supports colors at all?
- Can it show 256 colors?
- Can it show truecolor


## Features
- Detects terminal color capability (dumb, 16, 256, truecolor)
- Parses `--color` and `--no-color` flags in both GNU and POSIX styles.
- Checks if color has being explicitly enabled or disabled via flags or environment variables.
- Respects NO_COLOR and FORCE_COLOR environment variables.
- Falls back to auto-detection when no color flags are provided
- Detects if output is a TTY.
- Automatically handles when user uses `--color=auto` or `--color=tty` or `--color tty`

---

## Installation
```bash
go get -u github.com/phamio/termcolor
```

---

## Usage

### Check color support

```go
package main
import (
    "fmt"
    "github.com/phamio/termcolor"
)

func main(){
     //Check if terminal supports color
    termCol := termcolor.SupportsColor()
    if termCol {
        fmt.Println("Hurray! Terminal understand colors")
    }
}
```

---

### TTY detection

```go
package main
import (
    "fmt"
    "github.com/phamio/termcolor"
)

func main(){
     //Check if output is to terminal
    if termcolor.IsTerminal() {
        fmt.Println("I'm printed on the terminal")
    }
}
```

---

### Check color toggle (enabled/disabled)
```go
package main
import (
    "fmt"
    "github.com/phamio/termcolor"
)

func main(){
    //Check if terminal color is explicitly enabled via flags or environment variables
    if termcolor.Enabled() {
        fmt.Prinln("Terminal color is enabled")
    }   

    //Check if terminal color is explicitly disabled via flags or environment variables
    if termcolor.Disabled() {
        fmt.Prinln("Terminal color is disabled")
    }   

    
    else {
        //--color=auto used or no preference given
        //termcolor decides from terminal
        if term.SupportsColor(){
            fmt.Println("Terminal wants colored texts")
        } else{
        fmt.Prinln("Terminal wants no colored texts")
        }
    }   
}
```

---

### Check capability level
```go
package main
import (
    "fmt"
    "github.com/phamio/termcolor"
)

func main(){
    cap := termcolor.Capability()

    switch cap {
    case termcolor.ColorTrue:
        //16 million colors
        fmt.Println("Terminal supports true color")
    case termcolor.Color256:
        //256 color palette
        fmt.Println("Terminal supports 256-color palette")
    case termcolor.Color16:
       //Basic ansi colors
        fmt.Println("Terminal supports 16 colors")
    case termcolor.ColorNone:
        //No color, plain output
        fmt.Println("Terminal supports no color (dumb terminal)")
    }
}
```

---

## License
Apache License 2.0 — see [LICENSE](LICENSE) file for details.

## About
termcolor was originally built as the color detection and flag parsing foundation for [Inkstamp](https://github.com/inkstamp/inkstamp)