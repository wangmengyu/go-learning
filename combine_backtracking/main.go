package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	res := combine(11, 11)
	for _, row := range res {
		fmt.Println(row)
	}
	fmt.Println(len(res))
	elapsed := time.Since(start)
	fmt.Printf("took %s\n", elapsed)

}

func combine(n int, k int) [][]int {
	res := [][]int{}
	var backtrack func(int, []int)
	backtrack = func(idx int, nums []int) {
		if len(nums) == k {
			t := make([]int, k)
			copy(t, nums)
			res = append(res, t)
			return
		}
		for i := idx; i <= n; i++ {
			backtrack(i+1, append(nums, i))
		}
	}

	backtrack(1, []int{})
	return res
}
