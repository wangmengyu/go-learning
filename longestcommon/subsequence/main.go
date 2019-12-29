package main

import (
	"fmt"
	"math"
)

func main() {
	s1 := "FOSHXWESIFJIE"
	s2 := "FISHMWAAIFSSS"

	cells := make([][]int, len(s2))

	for i, a := range s2 {
		cells[i] = make([]int, len(s1))
		for j, b := range s1 {
			if a == b {
				cells[i][j] = 1
				if i-1 >= 0 && j-1 >= 0 {
					cells[i][j] = 1 + cells[i-1][j-1]
				}
			} else {
				leftVal := 0
				if j-1 >= 0 {
					leftVal = cells[i][j-1]
				}
				upVal := 0
				if i-1 >= 0 {
					upVal = cells[i-1][j]
				}
				cells[i][j] = int(math.Max(float64(leftVal), float64(upVal)))
			}
		}
	}

	maxVal := cells[len(s2)-1][len(s1)-1]
	fmt.Println(cells)
	fmt.Println("max val:", maxVal)

}
