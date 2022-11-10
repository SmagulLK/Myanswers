package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argLen := len(os.Args[1:])
	if argLen == 0 {
		z01.PrintRune('0')
	}
	str := itoa(argLen)
	for _, w := range str {
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}

func itoa(n int) string {
	var result string
	var oneRune rune
	for n != 0 {
		oneRune = rune(n%10) + 48
		result = string(oneRune) + result
		n = n / 10
	}
	return result
}
