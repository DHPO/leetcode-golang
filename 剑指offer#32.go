package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counter := make([]int, 26)
	for _, c := range []byte(s) {
		counter[c - 'a'] += 1
	}
	for _, c := range []byte(t) {
		counter[c - 'a'] -= 1
		if counter[c - 'a'] < 0 {
			return false
		}
	}
	for i := range []byte(s) {
		if []byte(s)[i] != []byte(t)[i] {
			return true
		}
	}
	return false
}

func main() {
	isAnagram("a", "a")
}