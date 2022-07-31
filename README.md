# Timeout

Timeout given function written by Go

# Usage

```go
package main

import (
	"fmt"
	"time"
	"github.com/igtm/timeout"
)

func main() {
    if err := timeout.Timeout(doSomething(), 3*time.Second); err != nil {
        if err == timeout.ErrTimeout {
            fmt.Printf("TIMEOUT: %s\n", err)
        } else {
            fmt.Println(err)
        }
    }
}
```