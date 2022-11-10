package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	num, _ := strconv.Atoi(args[0])
	if num <= 1 {
		return
	}
	Fprime(num)
	fmt.Println()
}

func Fprime(n int) {
	for i := 2; i < n*2; i++ {
		if n%i == 0 {
			fmt.Print(i)
			if n/i != 1 {
				fmt.Print("*")
			}
			Fprime(n / i)
			break
		}
	}
}
