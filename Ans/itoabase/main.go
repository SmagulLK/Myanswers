package main

func ItoaBase(value, base int) string {
	sign := ""
	if value < 0 {
		sign = "-"
		value *= -1
	}
	result := ""
	for value > 0 {
		rem := value % base
		if rem < 10 {
			result = itoa(rem) + result
		} else {
			result = string(rune(rem+'A'-10)) + result
		}
		value /= base
	}
	return sign + result
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	s := ""
	for n > 0 {
		s = string(rune(n%10+48)) + s
		n /= 10
	}
	return s
}
