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
		result := []rune(args[0])
		for i, letter := range result {
			if letter >= 'a' && letter <= 'z' {
				result[i] = 'a' + 'z' - letter
			} else if letter >= 'A' && letter <= 'Z' {
				result[i] = 'A' + 'Z' - letter
			}
		}
		for _, w := range string(result) {
			z01.PrintRune(w)
		}
		z01.PrintRune('\n')
	}
}

// func main() {
// 	s := make(map[rune]rune)
// 	for i, j := 'A', 'Z'; i <= 'Z'; i++ {
// 		s[i] = j
// 		j--
// 	}
// 	for i, j := 'a', 'z'; i <= 'z'; i++ {
// 		s[i] = j
// 		j--
// 	}
// 	arg := os.Args[1]
// 	for i := 0; i < len(arg); i++ {
// 		if v, ok := s[rune(arg[i])]; ok {
// 			z01.PrintRune(v)
// 		} else {
// 			z01.PrintRune(rune(arg[i]))
// 		}
// 	}
// }
