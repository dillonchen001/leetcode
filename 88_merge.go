package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1

	if n == 0 {
		return
	}

	ser := m + n - 1
	for i >= 0 && j >= 0 {
		if nums2[j] > nums1[i] {
			nums1[ser] = nums2[j]
			ser--
			j--
		} else {
			nums1[ser] = nums1[i]
			ser--
			i--
		}
	}

	for i >= 0 {
		nums1[ser] = nums1[i]
		ser--
		i--
	}

	for j >= 0 {
		nums1[ser] = nums2[j]
		ser--
		j--
	}
}
