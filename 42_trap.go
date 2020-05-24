package main

func trap(height []int) int {
	ln := len(height)
	sum := 0
	if ln < 3 {
		return sum
	}

	maxLeft := 0
	maxRight := 0
	left := 1
	right := ln - 2

	for i := 1; i < ln-1; i++ {
		if height[left-1] < height[right+1] {
			maxLeft = myMax(maxLeft, height[left-1])
			if maxLeft > height[left] {
				sum = sum + (maxLeft - height[left])
			}
			left++
		} else {
			maxRight = myMax(maxRight, height[right+1])
			if maxRight > height[right] {
				sum = sum + (maxRight - height[right])
			}
			right--
		}
	}
	return sum
}

func myMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
