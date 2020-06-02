package main

func removeDuplicates(nums []int) int {
	ln := len(nums)
	if ln <= 2 {
		return ln
	}

	repeat := 1
	k := 0
	for i := 1; i < ln; i++ {
		k++
		if nums[i] == nums[i-1] {
			repeat++
		} else {
			repeat = 1
		}

		if repeat == 3 {
			k--
			repeat--
		} else {
			nums[k] = nums[i]
		}

	}

	return k + 1

}
