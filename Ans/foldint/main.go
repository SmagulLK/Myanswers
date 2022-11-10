package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}

func FoldInt(f func(int, int) int, a []int, n int) {
	for i := 0; i < len(a); i++ {
		n = f(n, a[i])
	}
	for _, w := range Itoa(n) {
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}

func Itoa(n int) string {
	result := ""
	if n == 0 {
		z01.PrintRune('0')
	}
	for n != 0 {
		result = string(rune(n%10)+48) + result
		n = n / 10
	}
	return result
}

func Add(a int, b int) int {
	return a + b
}

func Mul(a int, b int) int {
	return a * b
}

func Sub(a int, b int) int {
	return a - b
}
