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
		for _, x := range args[0] {
			for z := 0; z < times(x); z++ {
				z01.PrintRune(x)
			}
		}
	}

	z01.PrintRune('\n')
}

func times(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r-'a') + 1
	} else if r >= 'A' && r <= 'Z' {
		return int(r-'A') + 1
	} else {
		return 1
	}
}

// func main() {
// 	args := os.Args[1:]
// 	if len(args) == 1 {
// 		for _, w := range args[0] {
// 			if w >= 'a' && w <= 'z' {
// 				for i := 'a'; i <= w; i++ {
// 					z01.PrintRune(w)
// 				}
// 			} else if w >= 'A' && w <= 'Z' {
// 				for i := 'A'; i <= w; i++ {
// 					z01.PrintRune(w)
// 				}
// 			} else {
// 				z01.PrintRune(w)
// 			}
// 		}
// 	}
// }
