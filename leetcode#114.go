package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func flattenHelper(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return root
	}
	left, right := root.Left, root.Right
	if left != nil {
		flattenLeft := flattenHelper(left)
		root.Left = nil
		root.Right = left
		root = flattenLeft
	}
	if right != nil {
		flattenRight := flattenHelper(right)
		root.Right = right
		root = flattenRight
	}
	return root
}

func flatten(root *TreeNode)  {
	flattenHelper(root)
}