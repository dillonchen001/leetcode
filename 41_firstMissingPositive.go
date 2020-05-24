package main

import "fmt"

func main() {
	nums := []int{3, 4, -1, 1}
	fmt.Println(nums)
	fmt.Println(firstMissingPositive(nums))
}

func firstMissingPositive(nums []int) int {
	ln := len(nums)
	for i := 0; i < ln; i++ {
		for nums[i] <= ln && nums[i] > 0 && nums[nums[i]-1] != nums[i] {
			nums = mySwap(nums, i, nums[i]-1)
			fmt.Println(nums)
		}
	}
	fmt.Println(nums)

	for i := 0; i < ln; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return ln + 1
}

func mySwap(nums []int, i int, j int) []int {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
	return nums
}
