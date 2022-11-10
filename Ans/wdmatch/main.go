package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	temp := 0
	result := ""
	if len(args) != 2 {
		return
	}
	for i := 0; i < len(args[0]); i++ {
		for j := temp; j < len(args[1]); j++ {
			if args[0][i] == args[1][j] {
				temp = j
				result += string(args[0][i])
				break
			}
		}
	}
	if len(result) == len(args[0]) {
		for _, letter := range result {
			z01.PrintRune(letter)
		}
	} else {
		return
	}
	z01.PrintRune('\n')
}
