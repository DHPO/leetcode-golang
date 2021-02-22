package main

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	results := []int{}
	for i := range root.Children {
		c := root.Children[i]
		results = append(results, postorder(c)...)
	}
	results = append(results, root.Val)
	return results
}

func main() {

}