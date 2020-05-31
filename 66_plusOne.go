package main

func plusOne(digits []int) []int {
	ld := len(digits)

	ret := make([]int, ld+1)

	curr := 1
	for i := ld - 1; i >= 0; i-- {
		ret[i+1] = (digits[i] + curr) % 10
		curr = (digits[i] + curr) / 10
	}

	if curr > 0 {
		ret[0] = curr
		return ret
	}

	return ret[1 : ld+1]
}
