package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	result := []rune{}
	if len(args) != 1 {
		return
	} else {
		for _, letter := range args[0] {
			if letter >= 'a' && letter <= 'z' {
				if letter+13 >= 'z' {
					result = append(result, letter+13-26)
				} else {
					result = append(result, letter+13)
				}
			} else if letter >= 'A' && letter <= 'Z' {
				if letter+13 >= 'Z' {
					result = append(result, letter+13-26)
				} else {
					result = append(result, letter+13)
				}
			} else {
				result = append(result, letter)
			}
		}
		for _, w := range string(result) {
			z01.PrintRune(w)
		}
		z01.PrintRune('\n')
	}
}
