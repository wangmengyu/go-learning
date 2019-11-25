package main

import "fmt"

func main() {
	myList := []int{19, 97, 9, 17, 1, 11, 10, 5, 3, 20}
	myList = quickSort(myList)
	fmt.Println(myList)

}

const LEFT = "left"
const RIGHT = "right"

/**
快速排序
*/
func quickSort(myList []int) []int {
	pivot := myList[0] // 设置左边第一个节点是基准， 最左边的位置已经腾出来了
	left := 0
	right := len(myList) - 1
	actPoint := RIGHT //设定开始移动的指针是右指针
	for left <= right {
		if left == right {
			//左有指针相同的时，将基准放入指针位置中
			myList[left] = pivot
			if left > 1 {
				fmt.Println("left qs:", myList)
				//递归处理左子序列
				quickSort(myList[0:left])
			}
			if right < len(myList)-2 {
				fmt.Println("right qs:", myList)
				//递归处理右子序列
				quickSort(myList[right+1:])
			}
			//返回处理好的序列
			return myList
		}

		if actPoint == RIGHT { //当前是右指针活动
			if myList[right] < pivot {
				//右指针所指的数字小于基准，把它的值放到左指针所在位置，并且交换移动权
				myList[left] = myList[right]
				actPoint = LEFT
				left++
			} else {
				//右指针所指的数字大于基准，继续前进
				right = right - 1
			}
		} else if actPoint == LEFT { //当前是左指针活动
			//如果左指针指向的数大于基准，则将该值放入右指针所指的位置，并且交换移动权限
			if myList[left] > pivot {
				myList[right] = myList[left]
				actPoint = RIGHT
				right = right - 1
			} else {
				//左指针所指的数字小于基准，继续前进
				left = left + 1
			}
		}
	}
	return myList
}
