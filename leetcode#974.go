package main

func subarraysDivByK(A []int, K int) int {
	PreSumCount := make(map[int]int)
	PreSum := 0
	result := 0

	for _, v := range A {
		PreSum = (PreSum + v + 5000 * K) % K
		result += PreSumCount[PreSum]
		PreSumCount[PreSum] += 1
	}
	if PreSum == 0 {
        result += 1
    }

	return result
}
