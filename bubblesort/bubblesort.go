package main

import "fmt"

func main() {
	queue := []int{5, 9, 3, 1, 2, 8, 4, 7, 6}
	queue = bubbleSort(queue)
	fmt.Println(queue)

}

/**
  冒泡排序
*/
func bubbleSort(queue []int) []int {
	totalLen := len(queue)
	for i := 0; i < totalLen-1; i++ {
		//从后向前依次访问相邻节点，如果左边大于右边，则进行交换
		p1 := totalLen - 2
		p2 := totalLen - 1
		for p1 >= i {
			if queue[p1] > queue[p2] {
				queue[p1], queue[p2] = queue[p2], queue[p1]
			}
			p1 = p1 - 1
			p2 = p2 - 1
		}
		//每次循环处理后会将最小节点换到当前比较序列最左边，固定下来，位置=i，再用同样方式排序好后续的节点
	}
	return queue
}
