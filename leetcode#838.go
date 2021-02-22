package main

import "fmt"

func pushDominoes(dominoes string) string {
	lDistance := make([]int, len(dominoes))
	rDistance := make([]int, len(dominoes))

	lastL := 100000
	for i := len(dominoes) - 1; i >= 0; i-- {
		if dominoes[i] == 'L' {
			lastL = 0
		} else if dominoes[i] == 'R' {
			lastL = 100000
		}		else{
			lastL ++
		}
		lDistance[i] = lastL
	}

	lastR := 100000
	result := make([]byte, len(dominoes))
	for i := 0; i < len(dominoes); i ++ {
		if dominoes[i] == 'R' {
			lastR = 0
		} else if dominoes[i] == 'L' {
			lastR = 100000
		} else {
			lastR ++
		}
		rDistance[i] = lastR
		l, r := lDistance[i], lastR
		if l > r && r < 100000 {
			result[i] = 'R'
		} else if l < r && l < 100000{
			result[i] = 'L'
		} else {
			result[i] = '.'
		}
	}

	return string(result)
}

func main() {
	fmt.Println(pushDominoes(".L.R...LR..L.."))
}