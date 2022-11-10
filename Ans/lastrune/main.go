package main

import (
	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(LastRune("♥tr♥"))
}

func LastRune(s string) rune {
	var result []rune
	for _, letter := range s {
		result = append(result, letter)
	}
	return result[len(result)-1]
}
