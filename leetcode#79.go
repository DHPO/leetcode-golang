package main

type Index struct {
	I int
	J int
}

func existAt(board [][]byte, visited map[Index]bool, word string, i, j int) bool {
	if _, ok := visited[Index{i, j}]; ok {
		return false
	}
	if len(word) == 0 {
		return true
	}
	if i < 0 || i >= len(board) {
		return false
	}
	if j < 0 || j >= len(board[0]) {
		return false
	}
	if board[i][j] != word[0] {
		return false
	}
	visited[Index{i, j}] = true
	if existAt(board, visited, word[1:], i + 1, j) || existAt(board, visited,word[1:], i - 1, j) || existAt(board, visited,word[1:], i, j + 1) || existAt(board, visited,word[1:], i, j - 1) {
		return true
	} else {
		delete(visited, Index{i, j})
		return false
	}
}

func exist(board [][]byte, word string) bool {
	for i, row := range board {
		for j, _ := range row {
			visited := make(map[Index]bool)
			if existAt(board, visited, word, i, j) {
				return true
			}
		}
	}
	return false
}
