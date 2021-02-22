package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	dim0, dim1 := len(matrix), len(matrix[0])
	dim0Min, dim0Max := 0, dim0-1
	dim1Min, dim1Max := 0, dim1-1
	for dim0Min < dim0Max || dim1Min < dim1Max {
		for ; dim1Max >= 0 && dim0Min <= dim0 - 1 && matrix[dim0Min][dim1Max] < target; dim0Min++ {
		}
		for ; dim1Min <= dim1 - 1 && dim0Max >= 0 && matrix[dim0Max][dim1Min] > target; dim0Max-- {
		}
		for ; dim0Max >= 0 && dim1Min <= dim1 - 1 && matrix[dim0Max][dim1Min] < target; dim1Min++ {
		}
		for ; dim0Min <= dim0 - 1 && dim1Max >= 0 && matrix[dim0Min][dim1Max] > target; dim1Max-- {
		}
        if dim0Max < 0 || dim1Max < 0 || dim0Min >= dim0 || dim1Min >= dim1 {
			return false
		}
        if matrix[dim0Min][dim1Max] == target {
            return true
        }
	}
	return matrix[dim0Min][dim1Max] == target
}

func main() {
	fmt.Println(searchMatrix([][]int{{1}}, 1))
}
