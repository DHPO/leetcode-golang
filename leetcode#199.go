package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	nextLevel := []*TreeNode{root}
	result := []int {}

	for {
		currLevel := nextLevel
		nextLevel = []*TreeNode{}
		if len(currLevel) == 0 {
			break
		}
		result = append(result, currLevel[0].Val)
		for _, n := range currLevel {
			if n.Right != nil {
				nextLevel = append(nextLevel, n.Right)
			}
			if n.Left != nil {
				nextLevel = append(nextLevel, n.Left)
			}
		}
	}

	return result
}
