package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 2, 3}
	fmt.Println(subsetsWithDup(nums))
}

func subsetsWithDup(nums []int) [][]int {
	var result [][]int
	result = append(result, []int{})
	ln := len(nums)
	if ln == 0 {
		return result
	}
	sort.Ints(nums)

	start := 0
	lastNum := 0
	//长度循环, 每层根据上次结果扩展
	for long := 1; long <= ln; long++ {
		lr := len(result)
		nStart := long - 1
		for i := start; i < lr; i++ {
			start++

			if long > 1 {
				lastNum = result[i][long-2]
			}

			for j := nStart; j < ln; j++ {
				if j+1 < ln && nums[j] == nums[j+1] {
					continue
				}

				if long > 1 {
					if lastNum > nums[j] {
						continue
					}

					if lastNum == nums[j] {
						if j > 0 && nums[j] > nums[j-1] {
							continue
						}

						//去调重复的
						m := 1
						valid := true
						for m <= long-1 {
							fmt.Println(long - 1)
							fmt.Println(m)
							if result[i][long-1-m] == lastNum {
								if nums[j-m] != lastNum {
									valid = false
									break
								} else {
									m++
									continue
								}
							} else {
								break
							}
						}

						if !valid {
							continue
						}
					}
				}

				var tmp []int
				for _, v := range result[i] {
					tmp = append(tmp, v)
				}
				tmp = append(tmp, nums[j])
				result = append(result, tmp)
			}
		}
	}
	return result
}
