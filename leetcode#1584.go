package main

import "sort"

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func distance(p1, p2 []int) int {
	return abs(p1[0] - p2[0]) + abs(p1[1] - p2[1])
}

func toRoot(parents []int, idx int) int {
	for parents[idx] != idx {
		idx = parents[idx]
	}
	return idx
}

func minCostConnectPoints(points [][]int) int {
	parents := make([]int, len(points))
	for i := range parents {
		parents[i] = i
	}
	distances := [][]int{}
	for i := 0; i < len(points) - 1; i ++ {
		p1 := points[i]
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]
			distances = append(distances, []int{i, j, distance(p1, p2)})
		}
	}
	sort.Slice(distances, func(i, j int) bool {return distances[i][2] < distances[j][2]})
	total := 0
	edges := 0
	for _, edge := range distances {
		i, j, d := edge[0], edge[1], edge[2]
		if p1, p2 := toRoot(parents, i), toRoot(parents, j); p1 != p2 {
			total += d
			edges ++
			if edges == len(points) - 1 {
				break
			}
			parents[p2] = p1
			parents[j] = p1
		}
	}
	return total
}
