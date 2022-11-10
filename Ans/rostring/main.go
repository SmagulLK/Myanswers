package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		z01.PrintRune('\n')
		return
	} else {
		for _, w := range rostring(args[0]) {
			z01.PrintRune(w)
		}
		z01.PrintRune('\n')
	}
}

func rostring(s string) string {
	a := split(s)
	result := ""
	for i := 1; i < len(a); i++ {
		result += string(a[i]) + " "
	}
	result += a[0]
	return result
}

func split(s string) []string {
	result := []string{}
	res := ""
	for i, w := range s {
		if w != ' ' {
			res += string(w)
		}
		if len(res) != 0 && (w == ' ' || i == len(s)-1) {
			result = append(result, res)
			res = ""
		}
	}
	return result
}
