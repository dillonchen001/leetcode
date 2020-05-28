package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)

	if (len1+len2)%2 == 1 {
		return int2float64(findKthNum(nums1, 0, len1, nums2, 0, len2, (len1+len2+1)/2))
	} else {
		return (int2float64(findKthNum(nums1, 0, len1, nums2, 0, len2, (len1+len2)/2)) + int2float64(findKthNum(nums1, 0, len1, nums2, 0, len2, (len1+len2)/2+1))) / 2
	}
}

func findKthNum(nums1 []int, start1, len1 int, nums2 []int, start2, len2, k int) int {
	if len1 > len2 {
		return findKthNum(nums2, start2, len2, nums1, start1, len1, k)
	}

	if len1 == 0 {
		return nums2[start2+k-1]
	}

	if k == 1 {
		if nums1[start1] < nums2[start2] {
			return nums1[start1]
		} else {
			return nums[start2]
		}
	}

	L1 := k / 2
	if L1 > len1 {
		L1 = len1
	}
	L2 := k - L1
	if nums1[start1+L1-1] < nums2[start2+L2-1] {
		return findKthNum(nums1, start1+L1, len1-L1, nums2, start2, len2, k-L1)
	} else if nums1[start1+L1-1] == nums2[start2+L2-1] {
		return nums1[start1+L1-1]
	} else {
		return findKthNum(nums1, start1, len1, nums2, start2+L2, len2-L2, k-L2)
	}
}

func int2float64(num int) float64 {
	str := strconv.Itoa(num)
	f, _ := strconv.ParseFloat(str, 64)
	return f
}
