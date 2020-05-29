package main

import "fmt"

func main() {
	nums := []int{1, 0, 2}
	fmt.Println(canJump(nums))
}

func canJump(nums []int) bool {
	ln := len(nums)
	if ln <= 1 {
		return true
	}

	if nums[0] == 0 {
		return false
	}

	for i := ln - 2; i > 0; i-- {
		if nums[i] != 0 {
			continue
		}

		j := i - 1
		for j >= 0 {
			if nums[j] != 0 {
				break
			}
			j--
		}

		fmt.Println(j)
		fmt.Println(i)
		for k := j; k >= 0; k-- {
			if nums[k] > i-k {
				i = k
				break
			}

			if k == 0 {
				return false
			}
		}
	}
	return true

}
