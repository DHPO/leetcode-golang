package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	result := []int{}
	printEdge := func (iMin, iMax, jMin, jMax int) {
		i, j := iMin, jMin
		for j <= jMax {
			result = append(result, matrix[i][j])
			j ++
		}
		i += 1
		for i <= iMax {
			result = append(result, matrix[i][j])
			i ++
		}
		j -= 1
		for j >= jMin {
			result = append(result, matrix[i][j])
			j --
		}
		i -= 1
		for i >= iMin {
			result = append(result, matrix[i][j])
			i --
		}
	}
	iMin, iMax, jMin, jMax := 0, len(matrix) - 1, 0, len(matrix[0]) - 1
	for iMin <= iMax && jMin <= jMax {
		printEdge(iMin, iMax, jMin, jMax)
		iMin += 1
		iMax -= 1
		jMin += 1
		jMax -=1
	}

	return result
}