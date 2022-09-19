# go-ionice
_OS independent ionice for Golang_

## Introduction

This repository provides a Golang package for setting process IO priority (like ionice does on Linux) in an
OS independent way. Currently Linux, MacOS and Windows are supported.

## Example

```go
package main 

import (
    "fmt"
    
    "github.com/secDre4mer/go-ionice"
)

func main() {
	niceness, _ := ionice.GetIoPriority()
	fmt.Println("IO priority is:", niceness)

	ionice.SetIoPriority(ionice.Low)
	fmt.Println("Changed IO priority to low!")
}
```

Error handling is omitted for brevity.