package main

import "fmt"

func main() {
	myList := []int{6, 2, 1, 7, 9, 3, 4, 8, 5}
	myList = selectSort(myList)
	fmt.Println(myList)

}

/**
选择排序
*/
func selectSort(myList []int) []int {
	//每次遍历要处理的序列
	for i := 0; i < len(myList); i++ {
		//从当前未处理队列中找出值最小的，与i位置的值进行交换
		for j := i + 1; j < len(myList); j++ {
			if myList[j] < myList[i] {
				myList[j], myList[i] = myList[i], myList[j]
			}
		}
		//一次循环完毕后，未处理序列最小的值会出现再i位置。
	}

	return myList
}
