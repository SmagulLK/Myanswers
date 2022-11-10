package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	} else {
		number, _ := strconv.Atoi(args[0])
		if ispowerof2(number) {
			result := "true"
			for _, w := range result {
				z01.PrintRune(w)
			}
			z01.PrintRune('\n')
		} else {
			result := "false"
			for _, w := range result {
				z01.PrintRune(w)
			}
			z01.PrintRune('\n')
		}
	}
}

func ispowerof2(n int) bool {
	return n&(n-1) == 0 && n > 0
}
