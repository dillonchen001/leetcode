package main

func search(nums []int, target int) bool {
	ln := len(nums)
	if ln == 0 {
		return false
	}

	start := 0
	end := ln - 1
	mid := 0
	for start <= end {
		mid = (start + end) / 2

		if nums[mid] == target {
			return true
		} else if nums[mid] < nums[end] {
			if nums[mid] < target && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		} else if nums[mid] > nums[end] {
			if target < nums[mid] && target >= nums[start] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			end--
		}
	}
	return false
}
