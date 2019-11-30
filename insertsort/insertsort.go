package main

import "fmt"

func main() {
	queue := []int{5, 3, 4, 7, 2, 8, 6, 9, 1}
	queue = insertSort(queue)
	fmt.Println(queue)
}

/**
插入排序
*/
func insertSort(queue []int) []int {
	for i := 0; i < len(queue)-1; i++ {
		//每次认为0-i的节点是已经处理过的序列。
		//获得i之后的一格节点
		//将J指向的节点分别与一处理的序列从右向左比较，发现左边大于右边的就交换，发现左边小于右边的就停止操作
		k := i + 1
		for j := i; j >= 0; j = j - 1 {
			if queue[j] > queue[k] {
				queue[j], queue[k] = queue[k], queue[j] // 值交换
				k = j                                   //处理数字指针 更新成交换后的位置
			} else {
				break
			}
		}
	}
	return queue

}
