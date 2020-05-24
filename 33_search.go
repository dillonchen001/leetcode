package main

func search(nums []int, target int) int {
	ln := len(nums)
	if ln < 1 {
		return -1
	}

	start := 0
	end := ln - 1
	mid := 0
	for start <= end {
		mid = (start + end) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[start] {
			if target >= nums[start] && target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if target > nums[mid] && target < nums[start] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}

	return -1
}
