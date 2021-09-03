package main

func corpFlightBookings(bookings [][]int, n int) []int {
	result := make([]int, n + 1)
	for _, booking := range bookings {
		start, end, count := booking[0], booking[1], booking[2]
		result[start - 1] += count
		result[end] -= count
	}
	current := 0
	for i, v := range result {
		result[i] = v + current
		current += v
	}
	return result[:n]
}
