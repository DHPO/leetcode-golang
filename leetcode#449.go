package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderEncode(root *TreeNode) string {
	var result strings.Builder
	result.WriteString(strconv.Itoa(root.Val) + ",")
	if root.Left != nil {
		result.WriteString(preorderEncode(root.Left))
	}
	if root.Right != nil {
		result.WriteString(preorderEncode(root.Right))
	}
	return result.String()
}

func toIntMap(s []string) []int {
	result := make([]int, len(s))
	for i, c := range s {
		result[i], _ = strconv.Atoi(c)
	}
	return result
}

func preorderDecode(data []int, idx int, limit int) (node *TreeNode, movedIdx int) {
	if idx >= len(data) {
		return
	}
	movedIdx = idx + 1
	node = &TreeNode{Val: data[idx]}
	if movedIdx < len(data) && data[movedIdx] < data[idx] {
		node.Left, movedIdx = preorderDecode(data, movedIdx, data[idx])
	}
	if movedIdx < len(data) && data[movedIdx] < limit {
		node.Right, movedIdx = preorderDecode(data, movedIdx, limit)
	}
	return
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	str := preorderEncode(root)
	return str[: len(str) - 1]
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	node, _ := preorderDecode(toIntMap(strings.Split(data, ",")), 0, 10001)
	return node
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
