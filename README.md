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

## Handling of OS specific details

Some OS (e.g. Linux) offer a more detailed IO priority interface, whereas others (e.g. MacOS) provide only 3 levels that
can be set. To abstract between these points of view, this package defines 5 niceness values (`VeryLow` to `VeryHigh`).

These niceness values are mapped to the OS specific values in some way: `SetIoPriority` is always valid with any niceness
value, but some OS implementations might map `High` and `VeryHigh` to the same underlying value. Some OS specific values
may also be mapped to the same `Niceness` despite being different.

Therefore, there is no guarantee that `GetIoPriority` returns the same value that was previously set with `SetIoPriority`,
and there is also no guarantee that calling `SetIoPriority` with the return value of `GetIoPriority` is a no-op.
This is typically true for the default values, but no promises beyond this are made.