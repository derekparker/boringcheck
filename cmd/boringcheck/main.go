package main

import (
	"fmt"
	"os"

	"github.com/derekparker/boringcheck/boringcheck"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: boringcheck [path]")
		os.Exit(1)
	}
	path := os.Args[1]

	fns := boringcheck.BoringCheck(path)
	for _, fn := range fns {
		fmt.Println(fn)
	}
}
