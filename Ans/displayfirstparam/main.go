package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1]
	for _, w := range args {
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}
