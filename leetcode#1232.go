package main

func checkStraightLine(coordinates [][]int) bool {
    if len(coordinates) <= 2 {
		return true
	}
	x, y := coordinates[1][0] - coordinates[0][0], coordinates[1][1] - coordinates[0][1]
	for i := 2; i < len(coordinates); i ++ {
		x_, y_ := coordinates[i][0] - coordinates[0][0], coordinates[i][1] - coordinates[0][1]
		if y_ * x - x_ * y != 0 {
			return false
		}
	}
	return true
}
