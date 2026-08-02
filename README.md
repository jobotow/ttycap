# TTYCap
**ttycap** tells you what colors your terminal can display and at what level or depth

**It answers questions like**
- Does my terminal supports colors at all?
- Can it show 256 colors?
- Can it show truecolor


## Features
- Detects terminal color capability (dumb, 16, 256, truecolor)
- Respects NO_COLOR and FORCE_COLOR environment variables.
- Detects if output is a TTY.

---

## Installation
```bash
go get -u codeberg.org/jobotow/ttycap
```

---

## Usage

### Check color support and tty

```go
package main
import (
    "fmt"
    "os"
    "codeberg.org/jobotow/ttycap"
)

func main(){
     //Check if terminal supports color
    cap := ttycap.Query(os.Stdout)
    if cap >= ttycap.ColorANSI16 {
        fmt.Println("Hurray! Terminal Understands colors")
    } else if cap != ttycap.IsTTY {
        fmt.Println("Output is piped(color disabled)")
    } else {
        fmt.Println("Nothing to see here")
    }
}
```

---

### Check capability level
```go
package main
import (
    "fmt"
    "os"
    "codeberg.org/jobotow/ttycap"
)

func main(){
    cap := ttycap.Query(os.Stdout)

    switch cap {
    case ttycap.ColorANSI16m:
        //16 million colors
        fmt.Println("Terminal supports true color")
    case ttycap.ColorANSI256:
        //256 color palette
        fmt.Println("Terminal supports 256-color palette")
    case ttycap.ColorANSI16:
       //Basic ansi colors
        fmt.Println("Terminal supports 16 colors")
    case ttycap.ColorNone:
        //No color, plain output
        fmt.Println("Terminal supports no color (dumb terminal)")
    }
}
```

---

## License
Apache License 2.0 - See [LICENSE](LICENSE) file for details 