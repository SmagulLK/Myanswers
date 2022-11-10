package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 || atoi(args[0]) <= 1 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
	} else {
		for _, w := range itoa(addprimesum(atoi(args[0]))) {
			z01.PrintRune(w)
		}
		z01.PrintRune('\n')
	}
}

func itoa(n int) string {
	res := ""
	for n != 0 {
		res = string(rune(n%10)+48) + res
		n = n / 10
	}
	return res
}

func addprimesum(n int) int {
	res := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			res += i
		}
	}
	return res
}

func isPrime(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func atoi(s string) int {
	res := 0
	temp := 1
	if s[0] == '-' {
		temp = -1
		s = s[1:]
	}
	for _, w := range s {
		if w >= '0' && w <= '9' {
			res = res*10 + int(w-48)
		}
	}
	return res * temp
}
