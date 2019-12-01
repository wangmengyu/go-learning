package main

import (
	"fmt"
)

/**
归并排序 核心操作
合并 [l,r] 之间的数据，
例如：
   []int{4,6,3,7}
   需要处理[4,6]与[3,7];
   两个指针分别指向两段数字的开头，逐个比较，将小的放在前面大的放在后面，一组组比较直到结束
*/
func merge(arr []int, l int, mid int, r int) {
	seq1 := arr[l : mid+1]
	seq2 := arr[mid+1 : r+1]
	fmt.Println("s1,s2", seq1, seq2)
	left := 0
	right := 0
	point := l
	//新建临时切片用于存储正确排序的合并序列
	tmpArr := make([]int, 0)
	for len(tmpArr) <= r-l {
		min := 0
		max := 0
		if left >= len(seq1) && right >= len(seq2) {
			//l 越界, r 越界，
			return
		} else if right >= len(seq2) && left < len(seq1) {
			//r 越界，l未越界
			tmpArr = append(tmpArr, seq1[left])
		} else if left >= len(seq1) && right < len(seq2) {
			//l 越界，r未越界
			tmpArr = append(tmpArr, seq2[right])
		} else {
			//r ,l 都不越界 ，进行比较，从小到大追加到临时数组中
			if seq1[left] > seq2[right] {
				min = seq2[right]
				max = seq1[left]
			} else {
				min = seq1[left]
				max = seq2[right]
			}
			tmpArr = append(tmpArr, min)
			tmpArr = append(tmpArr, max)
		}
		fmt.Println("min,max=", min, max)
		point = point + 2
		left++
		right++
	}
	for i, v := range tmpArr {
		arr[l+i] = v
	}
}

/**
递归的处理每一段数据，直到，两个指针指向同一个数字
*/
func MergeSort(arr []int, l int, r int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2
	MergeSort(arr, l, mid)
	MergeSort(arr, mid+1, r)
	merge(arr, l, mid, r)
}

func main() {
	arr := []int{4, 6, 3, 7, 5, 1, 2}
	MergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)

}
