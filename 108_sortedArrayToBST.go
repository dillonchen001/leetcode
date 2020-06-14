package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := (len(nums)+1)/2 - 1

	root := &TreeNode{nums[mid], nil, nil}

	if mid > 0 {
		fmt.Println(nums[0:mid])
		root.Left = sortedArrayToBST(nums[0:mid])
	}

	if mid+1 < len(nums) {
		fmt.Println(nums[mid+1:])
		root.Right = sortedArrayToBST(nums[mid+1:])
	}
	return root
}
