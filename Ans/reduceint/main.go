package main

import (
	"strconv"

	"github.com/01-edu/z01"
)

func ReduceInt(a []int, f func(int, int) int) {
	result := a[0]
	for i := 1; i < len(a); i++ {
		result = f(result, a[i])
	}
	res := strconv.Itoa(result)
	for _, w := range res {
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}
