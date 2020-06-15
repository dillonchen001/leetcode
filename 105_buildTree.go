package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	for j := 0; j < len(inorder); j++ {
		if inorder[j] == root.Val {
			root.Left = buildTree(preorder[1:j+1], inorder[0:j])
			root.Right = buildTree(preorder[j+1:], inorder[j+1:])
		}
	}
	return root

}
