package main

import "fmt"

func main() {
	fmt.Println(Itoa(955855))
}

func Itoa(n int) string {
	var res string
	var OneRune rune
	temp := ""
	if n == 0 {
		return "0"
	}
	if n < 0 {
		temp = "-"
		n = n * (-1)
		for n != 0 {
			OneRune = rune(n%10) + 48
			res = string(OneRune) + res
			n = n / 10
		}
		return temp + res
	}
	for n != 0 {
		OneRune = rune(n%10) + 48
		res = string(OneRune) + res
		n = n / 10
	}
	return res
}
