package main

func numPairsDivisibleBy60(time []int) int {
    count := make([]int, 60)
    result := 0
    for _, t := range time {
        count[t % 60] += 1
    }
    for t, c := range count[: 31] {
        if t == 0 || t == 30 {
            result += c * (c - 1) / 2
        } else {
            result += c * count[60 - t]
        }
    }
    return result
}