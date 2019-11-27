package main

import (
	"fmt"
	"math"
)

func maxSubStr(s1, s2 string) [][]int {
	cells := make([][]int, len(s1))
	for i, c1 := range s1 {
		cells[i] = make([]int, len(s2))
		for j, c2 := range s2 {
			lastVal := 0
			if i > 1 && j > 1 {
				// 斜上方的数字
				lastVal = cells[i-1][j-1]
			}
			if c1 == c2 {
				cells[i][j] = 1 + lastVal
			} else {
				cells[i][j] = 0
			}

		}
	}
	return cells
}

func main() {

	// 最长公共子串：如果不相等，记录0，如果相等 记录1+斜上方的值
	s1 := "HISH"
	s2 := "FISH"

	fmt.Println(maxSubStr(s1, s2))

	//最长公共子序列：获得左侧或者上部节点
	s3 := "clues"
	s4 := "blue"

	cells, res := maxSubSeq(s3, s4)

	fmt.Println(cells)
	fmt.Printf("%s", string(res))

}

/**
获得最大子序列
*/
func maxSubSeq(s string, s2 string) ([][]int, []byte) {
	cells := make([][]int, len(s))
	res := make([]byte, int(math.Max(float64(len(s)), float64(len(s2)))))
	for i, c1 := range s {
		cells[i] = make([]int, len(s2))
		for j, c2 := range s2 {
			//获得纵向最大值
			upVal := 0
			if i-1 >= 0 {
				upVal = cells[i-1][j]
			}
			//获得横向最大值
			leftVal := 0
			if j-1 >= 0 {
				leftVal = cells[i][j-1]
			}
			//获得横向与纵向的最大值
			leftTopMax := int(math.Max(float64(upVal), float64(leftVal)))
			if c1 != c2 {
				//若本次不相同，返回原有最大值
				cells[i][j] = leftTopMax
			} else {
				//如果本次相同，则在原有共同最大值上+1
				cells[i][j] = leftTopMax + 1
				fmt.Printf("%c\n", c1)
				res = append(res, byte(c1))
			}
		}
	}
	return cells, res

}
