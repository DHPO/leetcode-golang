package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a... int) int {
	result := -10000
	for _, v := range a {
		if v > result {
			result = v
		}
	}
	return result
}

var ans int = 0

func maxDownSinglePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDownSinglePath(root.Left)
	right := maxDownSinglePath(root.Right)

	larger := max(left, right)
	result := 0
	if root.Val + larger < 0 {
		result = 0
	} else {
		result = root.Val + larger 
	}
	path := max(root.Val, root.Val + left + right)
	if path > ans {
		ans = path
	}
	return result
}

func maxPathSum(root *TreeNode) int {
	ans = 0
	maxDownSinglePath(root)
	return ans
}
