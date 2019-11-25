package main

import "fmt"

func main() {
	myList := []int{4, 3, 8, 6}
	myList = selectSort(myList)
	fmt.Println(myList)

}

/**
选择排序。
从序列中，寻找最小的元素，挑出来，放入新的序列。
直到全部挑完
*/
func selectSort(myList []int) []int {
	var res []int
	for len(myList) > 0 {
		//寻找最小的元素 和他 的位置
		minVal, pos := findMin(myList)
		//前半部分和后半部分序列拼接起来 ， 除了找到的最小元素
		endList := myList[pos+1:]
		myList = myList[:pos]
		for _, v := range endList {
			myList = append(myList, v)
		}
		res = append(res, minVal)
	}

	return res

}

/**
查找序列中最小的元素值和他的位置
*/
func findMin(myList []int) (int, int) {
	minVal := myList[0]
	pos := 0
	for i, v := range myList {
		if v < minVal {
			minVal = v
			pos = i
		}
	}
	return minVal, pos
}
