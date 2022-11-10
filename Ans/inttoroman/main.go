package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 && !isNumber(args[0]) {
		fmt.Println("Error")
	} else {
		fmt.Println(formula(romantoint(args[0])))
		fmt.Println(operator(romantoint(args[0])))
	}
}

func isNumber(s string) bool {
	for _, w := range s {
		if w < 48 || w > 57 {
			return false
		}
	}
	return true
}

func romantoint(s string) []int {
	symbols := map[rune]int{'M': 1000, 'D': 500, 'C': 100, 'L': 50, 'X': 10, 'V': 5, 'I': 1}
	// length := len(s)
	nums := []int{}
	for _, v := range s {
		for key, element := range symbols {
			if v == key {
				nums = append(nums, element)
			}
		}
	}
	return nums
}

func operator(s []int) int {
	result := 0
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			result += s[i]
			break
		}
		if s[i] < s[i+1] {
			result += s[i+1] - s[i]
			i++
		} else {
			result += s[i]
		}
	}
	return result
}

func itoa(n int) string {
	var oneRune rune
	var result string
	for n != 0 {
		oneRune = rune(n%10) + 48
		result = string(oneRune) + result
		n = n / 10
	}
	return result
}

func formula(n []int) string {
	result := ""
	for i := 0; i < len(n); i++ {
		if i == len(n)-1 {
			result += itoa(n[len(n)-1])
			break
		}
		if n[i] < n[i+1] {
			result += "(" + itoa(n[i+1]) + " - " + itoa(n[i]) + ")" + " + "
			i++
		} else {
			result += itoa(n[i]) + " + "
		}
	}
	result = result[:len(result)]
	return result
}
