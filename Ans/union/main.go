package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		z01.PrintRune('\n')
	} else {
		s1 := args[0]
		s2 := args[1]
		result := ""
		for _, w := range s1 {
			if !contain(result, w) {
				result += string(w)
			}
		}
		for _, w := range s2 {
			if !contain(result, w) {
				result += string(w)
			}
		}
		for _, w := range result {
			z01.PrintRune(w)
		}
		z01.PrintRune('\n')
	}
}

func contain(s string, r rune) bool {
	for _, w := range s {
		if w == r {
			return true
		}
	}
	return false
}
