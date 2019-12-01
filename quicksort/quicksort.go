package main

import "fmt"

func main() {
	myList := []int{3, 5, 8, 1, 2, 9, 4, 7, 6}
	left := 0
	right := len(myList) - 2
	pivot := len(myList) - 1
	quickSort(myList, left, right, pivot)
	fmt.Println(myList)

}

const LEFT = "left"
const RIGHT = "right"

/**
快速排序
*/
func quickSort(myList []int, left, right, pivot int) {
	currentAct := LEFT
	for left < right {
		if currentAct == LEFT {
			if myList[left] < myList[pivot] {
				//左指针走到的数字小于基准数，继续前进
				left++
			} else {
				//否则交换行动权限给右指针
				currentAct = RIGHT
			}
		} else {
			//右指针移动中
			if myList[right] > myList[pivot] {
				right--
			} else {
				currentAct = LEFT
				//此时，两个指针都已经静止，需要交换
				myList[left], myList[right] = myList[right], myList[left]
			}
		}
	}

	if left == right {
		fmt.Println("left,right,pvote", myList[left], myList[pivot])
		myList[left], myList[pivot] = myList[pivot], myList[left]
		//处理左侧序列
		quickSort(myList, 0, left-2, left-1)
		//处理右侧序列
		quickSort(myList, left+1, pivot, pivot-1)
		return
	}

}
