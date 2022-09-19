package main

import (
	"fmt"
	"os"

	"github.com/secDre4mer/go-ionice"
)

func main() {
	niceness, err := ionice.GetIoPriority()
	if err != nil {
		fmt.Println("Could not get IO priority", err)
		os.Exit(1)
	}
	fmt.Println("IO priority is:", toString(niceness))

	if err := ionice.SetIoPriority(ionice.Low); err != nil {
		fmt.Println("Could not set IO priority", err)
		os.Exit(1)
	}
	fmt.Println("Successfully changed IO priority to low!")

	niceness, err = ionice.GetIoPriority()
	if err != nil {
		fmt.Println("Could not get IO priority", err)
		os.Exit(1)
	}
	fmt.Println("IO priority is now:", toString(niceness))
}

func toString(n ionice.Niceness) string {
	switch n {
	case ionice.VeryLow:
		return "very low"
	case ionice.Low:
		return "low"
	case ionice.Standard:
		return "standard"
	case ionice.High:
		return "high"
	case ionice.VeryHigh:
		return "very high"
	default:
		panic("bad niceness")
	}
}
