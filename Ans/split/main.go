package main

import "fmt"

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}

func Split(s, sep string) []string {
	slice := 0
	result := []string{}
	for i := range s {
		if i < len(s)-len(sep) {
			if s[i:i+len(sep)] == sep {
				result = append(result, s[slice:i])
				slice = i + len(sep)
			}
		}
		if i == len(s)-1 {
			result = append(result, s[slice:])
		}
	}
	return result
}
