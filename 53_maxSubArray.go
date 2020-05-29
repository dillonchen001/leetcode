package main

func maxSubArray(nums []int) int {
	ln := len(nums)
	if ln == 0 {
		return 0
	}

	total := 0
	max := nums[0]
	for i := 0; i < ln; i++ {
		total += nums[i]
		if total > max {
			max = total
		}

		if total < 0 {
			total = 0
		}
	}
	return max
}
