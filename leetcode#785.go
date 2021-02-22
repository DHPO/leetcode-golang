package main

import "fmt"

func dfs(graph *[][]int, node int, color *map[int]bool) bool {
	nodeColor := (*color)[node]
	for _, n := range (*graph)[node] {
		if c, ok := (*color)[n]; ok {
			if c == nodeColor {
				return false
			}
		} else {
			(*color)[n] = !nodeColor
			if !dfs(graph, n, color) {
				return false
			}
		}
	}
	return true
}

func isBipartite(graph [][]int) bool {
	color := make(map[int]bool)
	for i:=0; i < len(graph); i ++ {
		if _, ok := color[i]; !ok {
			if !dfs(&graph, i, &color) {
				return false
			}
		}
	}
	return true
}

func main () {
	fmt.Println(isBipartite(()))
}
