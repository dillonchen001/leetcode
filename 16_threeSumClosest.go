package main

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	result := 0
	length := len(nums)
	if length < 3 {
		return result
	}

	sort.Ints(nums)
	result = nums[0] + nums[1] + nums[2]

	for k := 0; k < length-2; k++ {
		if nums[k] > target && nums[k] > 0 {
			break
		}

		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		i := k + 1
		j := length - 1
		for i < j {
			if nums[k]+nums[i]+nums[j] < target {
				result = getMinDis(nums[k]+nums[i]+nums[j], result, target)
				i++
			} else if nums[k]+nums[i]+nums[j] > target {
				result = getMinDis(nums[k]+nums[i]+nums[j], result, target)
				j--
			} else {
				return target
			}
		}
	}
	return result
}

func getNumAbs(num int) int {
	if num >= 0 {
		return num
	}
	return 0 - num
}

func getMinDis(num1, num2, target int) int {
	if getNumAbs(num1-target) > getNumAbs(num2-target) {
		return num2
	}
	return num1
}
