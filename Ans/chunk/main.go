package main

import "fmt"

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

func Chunk(slice []int, size int) {
	arr := []int{}
	res := [][]int{}
	if size == 0 {
		fmt.Println()
	}
	for i := 0; i < len(slice); i++ {
		arr = append(arr, slice[i])
		if size == 0 {
			return
		}
		if len(arr) == size || i == len(slice)-1 {
			res = append(res, arr)
			arr = []int{}
		}
	}
	fmt.Println(res)
}
