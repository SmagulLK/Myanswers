package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	} else {
		result := ""
		for _, w := range args[0] {
			if string(w) == args[1] {
				result += args[2]
			} else {
				result += string(w)
			}
		}
		for _, w := range result {
			z01.PrintRune(w)
		}
	}
	z01.PrintRune('\n')
}
