package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	} else {
		number1, _ := strconv.Atoi(args[0])
		number2, _ := strconv.Atoi(args[1])
		if number1 > 0 && number2 > 0 {
			for number1 != number2 {
				if number1 > number2 {
					number1 = number1 - number2
				} else {
					number2 = number2 - number1
				}
			}
			fmt.Println(number1)
		}
	}
}
