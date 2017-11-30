package main

import (
	"os"
)

func main() {
	if len(os.Args) < 2 {
		Fprintf(os.Stderr, "%s: not enough arguments\n", []string{os.Args[0]})
	} else if len(os.Args) == 2 {
		f := os.Args[1]
		Printf(f, []string{})
		//Printf(f, []string{})
	} else {
		f := os.Args[1]
		s := os.Args[2:]
		Printf(f, s)
	}
}
