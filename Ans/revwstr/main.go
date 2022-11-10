package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	} else {
		for _, w := range revwstr(args[0]) {
			z01.PrintRune(w)
		}
		z01.PrintRune('\n')
	}
}

func revwstr(s string) string {
	a := split(s)
	result := ""
	for i := len(a) - 1; i >= 0; i-- {
		result += string(a[i]) + " "
	}
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
