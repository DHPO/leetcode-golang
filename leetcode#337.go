package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func rob(root *TreeNode) int {
	memA := make(map[*TreeNode]int)
	memB := make(map[*TreeNode]int)

	var robA func(root *TreeNode) int // not include self
	var robB func(root *TreeNode) int // include self
	robA = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		if v, ok := memA[root]; ok {
			return v
		}
		result := max(robA(root.Left), robB(root.Left)) + max(robA(root.Right), robB(root.Right))
		memA[root] = result
		return result
	}
	robB = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		if v, ok := memB[root]; ok {
			return v
		}
		result := robA(root.Left) + robA(root.Right) + root.Val
		memB[root] = result
		return result
	}

	return max(robA(root), robB(root))
}