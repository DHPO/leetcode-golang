package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	deps := make([][]int, numCourses)
	for _, dep := range prerequisites {
		from, to := dep[0], dep[1]
		if deps[from] == nil {
			deps[from] = []int{to}
		} else {
			deps[from] = append(deps[from], to)
		}
	}
	results := make([]int, 0, numCourses)
	visited := make([]byte, numCourses)
	err := false

	const (
		unvisited byte = iota
		pending
		added
	)

	var dfs func(n int)
	dfs = func(n int) {
		if err {
			return
		}
		if visited[n] == pending {
			err = true
		} else if visited[n] == added {
			return
		} else {
			visited[n] = pending
			if deps[n] != nil {
				for _, d := range deps[n] {
					dfs(d)
				}
			}
			visited[n] = added
			results = append(results, n)
		}
	}

	for n := 0; n < numCourses; n ++ {
		dfs(n)
	}

	if err {
		return []int{}
	}

	return results
}
