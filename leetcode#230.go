package main


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func dfs(node *TreeNode, c chan int) {
	if node != nil {
		dfs(node.Left, c)
		c <- node.Val
		dfs(node.Right, c)
	}
}

func kthSmallest(root *TreeNode, k int) int {
	c := make(chan int, 10)
	go dfs(root, c)
	count := 0

	for v := range c {
		count += 1
		if count == k {
			return v
		}
	}

	return 0
}