package main

import "fmt"

var paths = make(map[int][][]int)

func pathsToTarget(graph [][]int, curr int) [][]int {
	if path, ok := paths[curr]; ok {
		return path
	}
	next := graph[curr]
	currPath := [][]int{}
	for _, n := range next {
		if path, ok := paths[n]; ok {
			for _, p := range path {
				currPath = append(currPath, append([]int{curr}, p...))
			}
		} else {
			path = pathsToTarget(graph, n)
			for _, p := range path {
				currPath = append(currPath, append([]int{curr}, p...))
			}
		}
	}
	paths[curr] = currPath

	return  currPath
}

func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)
	paths[n-1] = [][]int{{n-1}}
	return pathsToTarget(graph, 0)
}

func main() {
	fmt.Println(allPathsSourceTarget([][]int{{4,3,1},{3,2,4},{3},{4},{}}))
}