package main

import "fmt"

func main() {
	str := "Hello world!"
	count := 0
	for i := 0; i < len(str); i++ {
		count++
	}
	fmt.Println(count)
}
