package main

import "fmt"

func main() {
	myList := []int{1, 3, 5, 7, 9}
	pos, ex := binary_search(myList, 9)
	fmt.Println(pos, ex)

}

/**
  二分查找
*/
func binary_search(myList []int, val int) (int, bool) {
	low := 0
	high := len(myList) - 1
	for low <= high {
		mid := (low + high) / 2
		midVal := myList[mid]
		fmt.Println(low, high, mid, midVal)
		if midVal == val {
			return mid, true
		}
		if val < midVal {
			//把高的指针用mid代替
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1, false

}
