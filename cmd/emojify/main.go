package main

import (
	"fmt"
	"os"

	"github.com/haukened/emojify"
)

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Fprintf(os.Stderr, "Error, only one command line argument allowed, in format :emoji:")
		os.Exit(1)
	} else {
		fmt.Println(emojify.Render(args[0]))
	}
}
