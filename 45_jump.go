package main

func jump(nums []int) int {
	ln := len(nums)
	index := 0
	steps := 0
	max := 0
	tmp := 0

	for index < ln-1 {
		steps++

		if index+nums[index] >= ln-1 {
			break
		}

		max = index + nums[index]
		tmp = index

		for i := tmp + 1; i <= tmp+nums[tmp]; i++ {
			if i+nums[i] > max {
				max = i + nums[i]
				index = i
			}
		}
	}
	return steps
}
