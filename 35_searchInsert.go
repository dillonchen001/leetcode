package main

func searchInsert(nums []int, target int) int {
	ln := len(nums)
	if ln < 1 {
		return 0
	}

	start := 0
	end := ln - 1
	mid := 0
	for start <= end {
		mid = (start + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	if nums[mid] < target {
		return mid + 1
	} else {
		return mid
	}
}
