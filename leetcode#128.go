package main

type Union struct {
	Data map[int]int
}

func (u *Union) GetRoot(e int) int {
	origin := e
	for u.Data[e] != e {
		e = u.Data[e]
	}
	u.Data[origin] = e
	return e
}

func (u *Union) Add(e int) {
	if _, ok := u.Data[e]; !ok {
		u.Data[e] = e
	}
}

func (u *Union) Connect(a, b int) {
	root1 := u.GetRoot(a)
	root2 := u.GetRoot(b)
	if root1 == root2 {
		return
	}
	u.Data[root2] = root1
}

func (u *Union) IsIn(a int) bool {
	_, ok := u.Data[a]
	return ok
}

func longestConsecutive(nums []int) int {
	union := Union{
		Data: make(map[int]int),
	}

	for _, n := range nums {
		union.Add(n)
		if union.IsIn(n - 1) {
			union.Connect(n - 1, n)
		}
		if union.IsIn(n + 1) {
			union.Connect(n + 1, n)
		}
	}

	count := make(map[int]int)
	dup := make(map[int]struct{})
	max := 0
	for _, n := range nums {
		if _, ok := dup[n]; ok {
			continue
		}
		dup[n] = struct{}{}
		root := union.GetRoot(n)
		count[root] += 1
		if count[root] > max {
			max = count[root]
		}
	}

	return max
}

func main() {
	longestConsecutive([]int{100,4,200,1,3,2})
}