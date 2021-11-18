package main

import "fmt"

// main 排列组合: 全排列就是第一个数字起每个数字分别与它后面的数字交换。
func main() {
	items := []int{1, 2, 3, 4, 5, 6, 7}
	num := 3
	resList := make([][]int, 0) // 收集全排列结果的二维数组
	AllRange(items, 0, num, &resList)
	fmt.Println(len(resList)) //
	fmt.Println(items)
}

func Swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

// 递归全排列， start 全排列开始的下标， length是数组长度
func AllRange(items []int, start int, length int, resList *[][]int) {
	if start == length-1 {
		// 已经到最后了
		itemCopy := make([]int, len(items))
		copy(itemCopy, items)
		*resList = append(*resList, itemCopy)
		fmt.Println(items)
		return
	}
	for i := start; i < length; i++ {
		//遍历开始位置， 分别与它后面的数字交换
		Swap(&items[start], &items[i])
		AllRange(items, start+1, length, resList)
		Swap(&items[start], &items[i])
	}

}
