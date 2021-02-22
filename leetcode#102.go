package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	curr := []*TreeNode{root}
	next := []*TreeNode{}
	result_layer := []int{}
	result := [][]int{}
	for {
		for _, h := range curr {
			if h != nil {
				result_layer = append(result_layer, h.Val)
				if h.Left != nil {
					next = append(next, h.Left)
				}
				if h.Right != nil {
					next = append(next, h.Right)
				}
			}
		}
		if len(result_layer) > 0 {
			result = append(result, result_layer)
		}
		result_layer = []int{}
		if len(next) == 0 {
			break
		}
		curr = next
		next = []*TreeNode{}
	}
	return result
}