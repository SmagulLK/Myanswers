package main

import "fmt"

func main() {
	a := []int{-123, 23, 1, 11, 55, -93}
	max := 0
	for i := 0; i < len(a); i++ {
		if max < a[i] {
			max = a[i]
		}
	}
	fmt.Println(max)
}
