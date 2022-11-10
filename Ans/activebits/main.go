package main

func ActiveBits(n int) int {
	var res []int
	count := 0
	for n > 0 {
		res = append(res, n%2)
		n = n / 2
	}
	for i := 0; i < len(res); i++ {
		if res[i] == 1 {
			count++
		}
	}
	return count
}
