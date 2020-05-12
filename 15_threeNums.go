package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var result [][]int
	length := len(nums)

	if length < 3 {
		return result
	}

	for k := 0; k < length-2; k++ {
		if nums[k] > 0 {
			break
		}

		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		i := k + 1
		j := length - 1
		target := 0 - nums[k]

		for i < j {
			if nums[i]+nums[j] > target {
				j--
			} else if nums[i]+nums[j] == target {
				tmp := []int{nums[k], nums[i], nums[j]}
				result = append(result, tmp)
				for i < j {
					if nums[i] == nums[i+1] {
						i++
					} else {
						break
					}
				}
				for i < j {
					if nums[j] == nums[j-1] {
						j--
					} else {
						break
					}
				}
				i++
				j--
			} else {
				i++
			}
		}
	}
	return result
}
