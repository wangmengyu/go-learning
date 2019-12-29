package main

import "fmt"

func main() {
	s1 := "HISX"
	s2 := "FISH"
	cells := make([][]int, len(s2))
	for i, a := range s2 {
		cells[i] = make([]int, len(s1))
		for j, b := range s1 {
			if a != b {
				cells[i][j] = 0
			} else {
				cells[i][j] = 1
				if i-1 >= 0 && j-1 >= 0 {
					cells[i][j] += cells[i-1][j-1]
				}
			}

		}
	}

	max := 0
	maxi := 0
	maxj := 0
	for i, _ := range s2 {
		for j, _ := range s1 {
			if cells[i][j] > max {
				max = cells[i][j]
				maxi = i
				maxj = j
			}
		}
	}
	fmt.Println(cells)
	fmt.Printf("longgest common subsequence = %d, i=%d,  j=%d", max, maxi, maxj)

}
