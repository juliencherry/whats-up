package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) >= 2 {
		green := "\033[0;32m"
		noColor := "\033[0m"
		fmt.Printf("%sAdded reminder: %s%s\n", green, noColor, args[1])
	} else {
		fmt.Println("What’s up?")
	}
}
