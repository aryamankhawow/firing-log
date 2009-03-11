package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}
	switch os.Args[1] {
	case "add":
		addCmd(os.Args[2:])
	case "grid":
		gridCmd(os.Args[2:])
	default:
		usage()
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "flog add|grid")
	os.Exit(1)
}

var _ = flag.String
