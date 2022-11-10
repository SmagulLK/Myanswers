package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) >= 1 {
		for _, w := range args {
			fmt.Println(brackets(w))
		}
	} else {
		fmt.Println("ERROR")
	}
}

func brackets(s string) string {
	str := []rune{}
	for _, w := range s {
		if w == '(' || w == '[' || w == '{' {
			str = append(str, w)
		}
		if w == ')' && str[len(str)-1] == '(' {
			str = str[:len(str)-1]
		}
		if w == ']' && str[len(str)-1] == '[' {
			str = str[:len(str)-1]
		}
		if w == '}' && str[len(str)-1] == '{' {
			str = str[:len(str)-1]
		}
	}
	if len(str) == 0 {
		return "OK"
	}
	return "ERROR"
}
