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
		tabmul(number)
	}
}

func tabmul(n int) {
	for i := 1; i < 10; i++ {
		result := itoa(i) + " x " + itoa(n) + " = " + itoa(i*n)
		for _, w := range result {
			z01.PrintRune(w)
		}
		z01.PrintRune('\n')
	}
}

func itoa(n int) string {
	result := ""
	for n != 0 {
		result = string(rune(n%10)+48) + result
		n = n / 10
	}
	return result
}
