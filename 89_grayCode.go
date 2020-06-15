package main

import "fmt"

func main() {
	fmt.Println(grayCode(2))
}

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}

	if n == 1 {
		return []int{0, 1}
	}

	last := grayCode(n - 1)
	ln := len(last)
	plus := 1 << uint(n-1)
	for i := ln - 1; i >= 0; i-- {
		last = append(last, last[i]+plus)
	}
	return last
}
