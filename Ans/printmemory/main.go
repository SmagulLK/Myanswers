package main

import (
	"unicode"

	"github.com/01-edu/z01"
)

func main() {
	a := [10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'}
	PrintMemory(a)
}

func PrintMemory(arr [10]byte) {
	hex := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F'}
	bnm := []string{}
	for _, w := range arr {
		if w == 0 {
			bnm = append(bnm, "00")
			continue
		}
		s := ""
		for w != 0 {
			s = string(hex[w%16]) + s
			w = w / 16
		}
		bnm = append(bnm, s)
	}
	for i := 0; i < len(bnm); i++ {
		if (i+1)%4 == 0 || i == len(bnm)-1 {
			console(bnm[i])
			z01.PrintRune('\n')
		} else {
			console(bnm[i] + " ")
		}
	}
	for _, w := range arr {
		if !unicode.IsGraphic(rune(w)) {
			z01.PrintRune('.')
		} else {
			z01.PrintRune(rune(w))
		}
	}
	z01.PrintRune('\n')
}

func console(s string) {
	for _, w := range s {
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}
