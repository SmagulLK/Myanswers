package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	arg := strings.Split(args[0], " ")
	operator := []string{}
	nums := []int{}
	fmt.Println(arg)
	for _, w := range arg {
		if w == "+" || w == "-" || w == "*" || w == "/" || w == "%" {
			operator = append(operator, string(w))
		} else if w >= "0" && w <= "9" {
			n, err := strconv.Atoi(w)
			if err != nil {
				return
			}
			nums = append(nums, n)
		} else if len(w) == 0 {
			continue
		} else {
			fmt.Println("Error")
			return
		}
	}

	result := 0
	temp := 0
	if len(nums) == len(operator)+1 {
		for i := 0; i < len(nums); i++ {
			for j := temp; j < len(operator); j++ {
				if i == len(nums)-1 {
					result = nums[i]
				} else if operator[j] == "+" {
					result = nums[i] + nums[i+1]
					nums[i+1] = result
					temp = j + 1
					break
				} else if operator[j] == "-" {
					result = nums[i] - nums[i+1]
					nums[i+1] = result
					temp = j + 1
					break
				} else if operator[j] == "*" {
					result = nums[i] * nums[i+1]
					nums[i+1] = result
					temp = j + 1
					break
				} else if operator[j] == "/" {
					result = nums[i] / nums[i+1]
					nums[i+1] = result
					temp = j + 1
					break
				} else if operator[j] == "%" {
					result = nums[i] % nums[i+1]
					nums[i+1] = result
					temp = j + 1
					break
				}
			}
		}
	} else {
		fmt.Println("ERROR")
		return
	}
	fmt.Println(result)
}
