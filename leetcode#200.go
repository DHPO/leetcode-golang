package main

type Index struct {
	I int
	J int
}

type Union struct {
	data map[Index]Index
	RootCount int
}

func (u *Union) New(index Index) {
	u.data[index] = index
	u.RootCount += 1
}

func (u *Union) GetRoot(index Index) Index {
	input := index
	for {
		if parent, ok := u.data[index]; ok && parent != index {
			index = parent
		} else {
			break
		}
	}
	u.data[input] = index
	return index
}

func (u *Union) Connect(idx1, idx2 Index) {
	root1, root2 := u.GetRoot(idx1), u.GetRoot(idx2)
	if root1 != root2 {
		u.data[root1] = root2
		u.RootCount -= 1
	}
}

func numIslands(grid [][]byte) int {
	u := Union{
		data: make(map[Index]Index),
		RootCount: 0,
	}
	for i, row := range grid {
		for j, _ := range row {
			if grid[i][j] == '1'{
				u.New(Index{i, j})
				if i > 0 && grid[i-1][j] == '1' {
					u.Connect(Index{i, j}, Index{i - 1, j})
				}
				if j > 0 && grid[i][j-1] == '1' {
					u.Connect(Index{i, j}, Index{i, j - 1})
				}
			}
		}
	}
	return u.RootCount
}
