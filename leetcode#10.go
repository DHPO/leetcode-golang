package main

import "fmt"

type twoMap struct {
	m map[int]map[int]bool
}

func (m *twoMap) Get(i, j int) (bool, bool) {
	if _, ok := m.m[i]; ok {
		v, ok := m.m[i][j]
		return v, ok
	}
	return false, false
}

func (m *twoMap) Set(i, j int, data bool) {
	if _, ok := m.m[i]; !ok {
		m.m[i] = make(map[int]bool)
	}
	m.m[i][j] = data
}

func isMatch(s string, p string) bool {
	buf := twoMap{m: map[int]map[int]bool{}}
	var f func(i, j int) bool
	f = func(i, j int) bool {
		if i < 0 && j < 0 {
			return true
		} else if i >= 0 && j < 0 {
			return false
		}
		if v, ok := buf.Get(i, j); ok {
			return v
		}
		if p[j] != '*' {
			if i < 0 {
				return false
			} else if p[j] == '.' {
				buf.Set(i, j, f(i-1, j-1))
			} else {
				buf.Set(i, j, s[i] == p[j] && f(i-1, j-1))
			}
		} else {
			if j == 0 {
				panic("Bad pattern")
			}
			if i < 0 {
				buf.Set(i, j, f(i, j-2))
			} else if p[j-1] != '.' {
				if s[i] != p[j-1] {
					buf.Set(i, j, f(i, j-2))
				} else {
					buf.Set(i, j, f(i-1, j) || f(i, j-2))
				}
			} else {
				buf.Set(i, j, f(i-1, j) || f(i, j-2))
			}
		}
		result, _ := buf.Get(i, j)
		return result
	}
	return f(len(s) - 1, len(p) - 1)
}

func main() {
	fmt.Println(isMatch("aaa", "aaaa"))
}
