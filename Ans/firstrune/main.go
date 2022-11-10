package main

import "github.com/01-edu/z01"

func main() {
	z01.PrintRune(FirstRune("♥tr"))
}

func FirstRune(s string) rune {
	var result []rune
	for i, letter := range s {
		if i == 0 {
			result = append(result, letter)
		}
	}
	return result[0]
}
