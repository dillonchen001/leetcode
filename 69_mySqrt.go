package main

import "fmt"

func main() {
	x := 8
	fmt.Println(mySqrt(x))
}

func mySqrt(x int) int {
	if x < 1 {
		return x
	}

	left := 1
	right := x / 2
	mid := 0
	for left < right {
		mid = (left + right + 1) / 2
		if mid*mid > x {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return left
}
