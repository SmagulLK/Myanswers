package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) == 1 {
		n, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("00000000")
		} else {
			fmt.Println(toBin(n))
		}
	}
}

func toBin(n int) string {
	res := ""

	for n > 0 {
		res = string(rune(n%2+48)) + res
		n = n / 2
	}
	for len(res) < 8 {
		res = "0" + res
	}
	return res
}
