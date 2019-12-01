package main

import "fmt"

/**
堆排序
一开始需要将n 个数据存进堆里，所需时间为O(nlogn)。排序过程中，堆从
空堆的状态开始，逐渐被数据填满。由于堆的高度小于log2n，所以插入1 个数据所需要
的时间为O(logn)。
每轮取出最大的数据并重构堆所需要的时间为O(logn)。由于总共有n 轮，所以重
构后排序的时间也是O(nlogn)。因此，整体来看堆排序的时间复杂度为O(nlogn)。
这样来看，堆排序的运行时间比之前讲到的冒泡排序、选择排序、插入排序的时间
O(n2) 都要短，但由于要使用堆这个相对复杂的数据结构，所以实现起来也较为困难。
*/
func main() {
	data := []int{1, 2, 4, 7, 3, 6, 11}
	for i, _ := range data {
		data = heapify(data, i)
	}
	seq := make([]int, 0)
	//循环取出第一个节点推送到到返回序列，并且将最后一格节点补充到0号位置，重新进行堆排序，知道所有节点处理完毕
	for len(data) > 0 {
		seq = append(seq, data[0])
		data[0] = data[len(data)-1]
		data = data[0 : len(data)-1]
		if len(data) == 0 {
			break
		}
		for j, _ := range data {
			data = heapify(data, j)
		}
		fmt.Println(data)
	}
	fmt.Println(seq)

}

/**
调整最小一颗数，保证第一个节点大于后续两个节点
*/
func heapify(data []int, parent int) []int {
	c1 := 2*parent + 1
	c2 := 2*parent + 2
	maxPos := len(data) - 1
	if parent < 0 {
		//已经处理完所有节点
		return data
	}
	changeFlag := 0
	if c1 <= maxPos && data[c1] > data[parent] {
		//左节点大于父节点，交换
		data[c1], data[parent] = data[parent], data[c1]
		changeFlag = 1
	}
	if c2 <= maxPos && data[c2] > data[parent] {
		//右节点大于父节点，交换
		data[c2], data[parent] = data[parent], data[c2]
		changeFlag = 2
	}

	parentParent := (parent - 1) / 2
	if changeFlag != 0 {
		//有过交换行为，需要处理当前节点的父节点树结构
		//查找当前节点的父节点位置，对它的父节点开头的最小树做堆排序
		if parent%2 == 0 {
			parentParent = (parent - 2) / 2
		}
		data = heapify(data, parentParent)
	}

	return data

}
