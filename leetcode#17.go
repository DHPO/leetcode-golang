package main

var digitMap = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 { return []string{}}
	if len(digits) == 1 {
		result := []string{}
		for _, c := range digitMap[digits[0]] {
			result = append(result, string(c))
		}
		return result
	}
	latterStrings := letterCombinations(digits[1:])
	results := []string{}
	for _, c := range digitMap[digits[0]] {
		for _, s := range latterStrings {
			results = append(results, string(c) + s)
		}
	}
	return results
}
