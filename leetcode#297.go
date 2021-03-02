package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


type Codec struct {
    
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    result := []string{}
	nextLevel := []*TreeNode{root}
	nextLevelNotNull := true

	for len(nextLevel) > 0 {
		currLevel := nextLevel
		nextLevel = []*TreeNode{}
		nextLevelNotNull = false

		for _, n := range currLevel {
			if n == nil {
				result = append(result, "null")
			} else {
				result = append(result, strconv.Itoa(n.Val))
				nextLevel = append(nextLevel, n.Left, n.Right)
				nextLevelNotNull = nextLevelNotNull || n.Left != nil || n.Right != nil
			}
		}

		if !nextLevelNotNull {
			break
		}
	}

	return strings.Join(result, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    nodes := strings.Split(data, ",")
	nodeP := 0
	var root *TreeNode = nil
	nextLevel := []**TreeNode{&root}

	for nodeP < len(nodes) {
		currLevel := nextLevel
		nextLevel = []**TreeNode{}
		for _, slot := range currLevel {
			value, err := strconv.Atoi(nodes[nodeP])
			nodeP ++
			var node *TreeNode
			if err != nil {
				node = nil
			} else {
				node = &TreeNode{
					Val: value,
				}
				nextLevel = append(nextLevel, &node.Left, &node.Right)
			}
			*slot = node
		}
	}
	return root
}