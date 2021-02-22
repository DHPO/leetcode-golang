package main

var buffer = map[int][]string {}

func generateParenthesis(n int) []string {
	if n == 0 {return []string{""}}
	if n == 1 {return []string{"()"}}
	if r, ok := buffer[n]; ok {
		return r
	}
	results := []string{}
	for i := 0; i < n ; i ++ {
		left := generateParenthesis(i)
		right := generateParenthesis(n-i-1)
		for _, l := range left {
			for _, r := range right {
				results = append(results, "(" + l + ")" + r)
			}
		}
	}
	buffer[n] = results
	return results
}