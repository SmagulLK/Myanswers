package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	} else {
		result := ""
		temp := 0
		for i := 0; i < len(args[0]); i++ {
			for j := temp; j < len(args[1]); j++ {
				if args[0][i] == args[1][j] {
					temp = j + 1
					result += string(args[0][i])
					break
				}
			}
		}
		if len(args[0]) == len(result) {
			z01.PrintRune('1')
		} else {
			z01.PrintRune('0')
		}
		z01.PrintRune('\n')
	}
}
