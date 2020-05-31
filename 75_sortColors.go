package main

func sortColors(nums []int) {
	ln := len(nums)
	if ln == 0 {
		return
	}

	p0 := 0
	start := 0
	end := ln - 1

	for start <= end {
		if nums[start] == 0 {
			nums[p0], nums[start] = nums[start], nums[p0]
			p0++
			start++
		} else if nums[start] == 2 {
			nums[end], nums[start] = nums[start], nums[end]
			end--
		} else {
			start++
		}
	}
}
