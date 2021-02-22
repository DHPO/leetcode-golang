package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	} else if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	root := &TreeNode{Val: preorder[0]}
	split := 0
	for i := range inorder {
		if inorder[i] == root.Val {
			split = i
			break
		}
	}
	root.Left = buildTree(preorder[1:split+1], inorder[0: split])
	root.Right = buildTree(preorder[split+1:], inorder[split+1:])
	return root
}
