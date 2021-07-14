//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
//
//
//
// 示例 1：
//
//
//输入：nums1 = [1,3], nums2 = [2]
//输出：2.00000
//解释：合并数组 = [1,2,3] ，中位数 2
//
//
// 示例 2：
//
//
//输入：nums1 = [1,2], nums2 = [3,4]
//输出：2.50000
//解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
//
//
// 示例 3：
//
//
//输入：nums1 = [0,0], nums2 = [0,0]
//输出：0.00000
//
//
// 示例 4：
//
//
//输入：nums1 = [], nums2 = [1]
//输出：1.00000
//
//
// 示例 5：
//
//
//输入：nums1 = [2], nums2 = []
//输出：2.00000
//
//
//
//
// 提示：
//
//
// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -106 <= nums1[i], nums2[i] <= 106
//
//
//
//
// 进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？
// Related Topics 数组 二分查找 分治算法
// 👍 4148 👎 0

package test

import (
	"fmt"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{2, 7, 11, 15}
	nums2 := []int{8, 9}
	res := findMedianSortedArrays(nums1, nums2)
	fmt.Println(res)
}

//leetcode submit region begin(Prohibit modification and deletion)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	l := (m + n + 1) / 2
	r := (m + n + 2) / 2
	return (float64(getKth(nums1, 0, nums2, 0, l)) + float64(getKth(nums1, 0, nums2, 0, r))) / 2.0
}

//leetcode submit region end(Prohibit modification and deletion)

func getKth(nums1 []int, s1 int, nums2 []int, s2, k int) int {
	if s1 > len(nums1)-1 {
		return nums2[s2+k-1]
	}

	if s2 > len(nums2)-1 {
		return nums1[s1+k-1]
	}

	// 特征情况(2)，分析见正文部分
	if k == 1 {
		if nums1[s1] >= nums2[s2] {
			return nums2[s2]
		} else {
			return nums1[s1]
		}
	}

	// 分别在两个数组中查找第k/2个元素，若存在（即数组没有越界），标记为找到的值；若不存在，标记为整数最大值
	mid1 := s1 + k/2 - 1
	mid2 := s2 + k/2 - 1
	if mid1 >= len(nums1) {
		mid1 = len(nums1) - 1
	}

	if mid2 >= len(nums2) {
		mid2 = len(nums2) - 1
	}
	if nums1[mid1] >= nums2[mid2] {
		return getKth(nums1, s1, nums2, s2+k/2, k-k/2)
	}

	return getKth(nums1, s1+k/2, nums2, s2, k-k/2)
}
