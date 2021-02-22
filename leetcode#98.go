package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isValidBSTRange(root *TreeNode, min, max int64) bool {
	if root == nil {
		return true
	}
	val := int64(root.Val)
	if val <= min || val >= max {
		return false
	}
	return isValidBSTRange(root.Left, min, val) && isValidBSTRange(root.Right, val, max)
}

func isValidBST(root *TreeNode) bool {
	return isValidBSTRange(root, -2147483649, 2147483648)
}