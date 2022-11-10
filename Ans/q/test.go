package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args[1:]) != 1 {
		return
	}
	var arry [2048]byte
	point := 0
	count := -1
	arg := os.Args[1]
	var arr []int
	for i := 0; i < len(arg); i++ {
		if arg[i] == '[' {
			arr = append(arr, i)
		}
	}
	for i := 0; i < len(arg); i++ {
		switch arg[i] {
		case '>':
			point++
		case '<':
			point--
		case '+':
			arr[point]++
		case '-':
			arr[point]--
		case '.':
			z01.PrintRune(rune(arr[point]))
		case '[':
			count++

		case ']':
			if arry[point] == 0 {
				count--
			} else {
				for j := i; j > 0; j-- {
					if arg[j] == '[' && arr[point] == j {
						i = j
					}
				}
			}

		}
	}
}
