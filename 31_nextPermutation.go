package main

import "fmt"

func main() {
	nums := []int{1, 3, 2}
	nextPermutation(nums)
	fmt.Println(nums)
}

func nextPermutation(nums []int) {
	ln := len(nums)

	if ln < 2 {
		return
	}

	i := ln - 1
	for i = ln - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			break
		}
	}

	if i == 0 {
		resort(nums, 0, ln-1)
		return
	}

	j := ln - 1
	for j = ln - 1; j > i-1; j-- {
		if nums[j] > nums[i-1] {
			break
		}
	}

	tmp := nums[j]
	nums[j] = nums[i-1]
	nums[i-1] = tmp

	resort(nums, i, ln-1)
}

func resort(nums []int, start, end int) {
	for start < end {
		a := nums[start]
		nums[start] = nums[end]
		nums[end] = a
		start++
		end--
	}
}
