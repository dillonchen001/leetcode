package main

func twoSum(nums []int, target int) []int {
	var result []int
	ln := len(nums)
	if ln < 2 {
		return result
	}

	for i := 0; i < ln-1; i++ {
		for j := i + 1; j < ln; j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i, j)
				return result
			}
		}
	}
	return result
}
