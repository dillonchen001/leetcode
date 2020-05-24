package main

func searchRange(nums []int, target int) []int {
	ln := len(nums)

	start := 0
	end := ln - 1
	mid := 0
	for start <= end {
		mid = (start + end) / 2
		if nums[mid] == target {
			return []int{searchMinLoc(nums, start, mid, target), searchMaxLoc(nums, mid, end, target)}
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return []int{-1, -1}
}

func searchMinLoc(nums []int, start, center, target int) int {
	mid := 0
	result := center
	for start <= center {
		mid = (start + center) / 2

		if nums[mid] < target {
			start = mid + 1
		} else {
			center = mid - 1
			result = mid
		}
	}
	return result
}

func searchMaxLoc(nums []int, center, end, target int) int {
	mid := 0
	result := center
	for center <= end {
		mid = (end + center) / 2

		if nums[mid] > target {
			end = mid - 1
		} else {
			center = mid + 1
			result = mid
		}
	}
	return result
}
