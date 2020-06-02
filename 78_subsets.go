package main

import "fmt"
import "sort"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}

func subsets(nums []int) [][]int {
	ln := len(nums)
	sort.Ints(nums)

	var ret [][]int

	ret = append(ret, []int{})
	start := 0

	for i := 0; i < ln; i++ {
		tmp := []int{nums[i]}
		ret = append(ret, tmp)
	}

	start = 1
	for i := 2; i <= ln; i++ {
		total := len(ret)
		for j := start; j < total; j++ {
			for k := i - 1; k < ln; k++ {
				if nums[k] <= ret[j][i-2] {
					continue
				}

				var tmp []int
				for _, v := range ret[j] {
					tmp = append(tmp, v)
				}
				tmp = append(tmp, nums[k])
				ret = append(ret, tmp)
			}
		}
		start = total
	}

	return ret
}
