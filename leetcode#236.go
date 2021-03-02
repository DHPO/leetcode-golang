package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p == q {
		return p
	}
	pPath := []*TreeNode{}
	qPath := []*TreeNode{}

	var search func(root, node *TreeNode, path *[]*TreeNode) bool
	search = func(root, node *TreeNode, path *[]*TreeNode) bool {
		if root == node {
			*path = append(*path, root)
			return true
		}
		if (root.Left != nil && search(root.Left, node, path)) || (root.Right != nil && search(root.Right, node, path)) {
			*path = append(*path, root)
			return true
		}
		return false
	}
	search(root, p, &pPath)
	search(root, q, &qPath)
	if len(pPath) < len(qPath) {
		pPath, qPath = qPath, pPath
	}
	for i := 0; i < len(qPath) ; i ++ {
		if pPath[len(pPath) - 1 - i] != qPath[len(qPath) - 1 - i] {
			return pPath[len(pPath) - i]
		}
	}
	return qPath[0]
}
