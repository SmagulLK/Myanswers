package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	s := ""
	str := args[0]
	for _, w := range str {
		if w != 'U' || w != 'I' || w != 'A' || w != 'O' || w != 'E' || w != 'u' || w != 'i' || w != 'a' || w != 'o' || w != 'e' {
			s = "No vowels"
		}
	}
	for i := range str {
		if str[i] == 'U' || str[i] == 'I' || str[i] == 'A' || str[i] == 'O' || str[i] == 'E' || str[i] == 'u' || str[i] == 'i' || str[i] == 'a' || str[i] == 'o' || str[i] == 'e' {
			s = str[i:] + str[:i] + "ay"
		}
	}
	fmt.Println(s)
}
