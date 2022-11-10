package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	} else if Atoi(args[0]) >= 4000 || !isNumber(args[0]) {
		err := "ERROR: cannot convert to roman digit"
		fmt.Printf("%v\n", err)
	} else {
		result := intToRoman(Atoi(args[0]))
		fmt.Printf("%v\n", result)
	}
}

func Atoi(s string) int {
	number := 0
	for _, w := range s {
		number = number*10 + int(w-48)
	}
	return number
}

func isNumber(s string) bool {
	for _, w := range s {
		if w < '0' || w > '9' {
			return false
		}
	}
	return true
}

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	result := ""
	formula := ""
	i := 0
	for num > 0 {
		for values[i] > num {
			i++
		}
		num -= values[i]
		result += romans[i]
		if len(romans[i]) == 2 {
			formula += "(" + string(romans[i][1]) + "-" + string(romans[i][0]) + ")" + "+"
		} else {
			formula += string(romans[i]) + "+"
		}
	}
	formula = formula[:len(formula)-1]
	fmt.Printf("%v\n", formula)
	return result
}
