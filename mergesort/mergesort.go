package main

import "fmt"

/**
归并排序中，分割序列所花费的时间不算在运行时间内（可以当作序列本来就是分
割好的）。在合并两个已排好序的子序列时，只需重复比较首位数据的大小，然后移动较
小的数据，因此只需花费和两个子序列的长度相应的运行时间。也就是说，完成一行归
并所需的运行时间取决于这一行的数据量。
看一下上面的图便能得知，无论哪一行都是n 个数据，所以每行的运行时间都为O(n)。
而将长度为n 的序列对半分割直到只有一个数据为止时，可以分成log2n 行，因此，总
共有log2n 行。也就是说，总的运行时间为O(nlogn)，这与前面讲到的堆排序相同。
*/
// 合并 [l,r] 两部分数据，mid 左半部分的终点，mid + 1 是右半部分的起点
func merge(arr []int, l int, mid int, r int) {
	// 因为需要直接修改 arr 数据，这里首先复制 [l,r] 的数据到新的数组中，用于赋值操作
	temp := make([]int, r-l+1)
	for i := l; i <= r; i++ {
		temp[i-l] = arr[i]
	}

	// 指向两部分起点
	left := l
	right := mid + 1

	for i := l; i <= r; i++ {
		// 左边的点超过中点，说明只剩右边的数据
		if left > mid {
			arr[i] = temp[right-l]
			right++
			// 右边的数据超过终点，说明只剩左边的数据
		} else if right > r {
			arr[i] = temp[left-l]
			left++
			// 左边的数据大于右边的数据，选小的数字
		} else if temp[left-l] > temp[right-l] {
			arr[i] = temp[right-l]
			right++
		} else {
			arr[i] = temp[left-l]
			left++
		}
	}
}

func MergeSort(arr []int, l int, r int) {
	if l >= r {
		return
	}

	// 递归向下
	mid := (r + l) / 2
	MergeSort(arr, l, mid)
	MergeSort(arr, mid+1, r)
	// 归并向上
	merge(arr, l, mid, r)
}

func main() {
	arr := []int{2, 1, 4, 7, 6, 3, 5}
	MergeSort(arr, 0, len(arr)-1)

	fmt.Println(arr)
}
