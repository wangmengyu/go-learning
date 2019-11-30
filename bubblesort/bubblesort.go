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
	p1 := len(queue) - 1
	p2 := len(queue) - 2
	res := make([]int, 0) //存储结果切片
	fmt.Println("p1,p2 = ", p1, p2)
	for len(queue) > 0 { //检查是否要处理的数据还有
		//定义最右侧的两个指针，分别指向最后两个元素
		p1 := len(queue) - 1
		p2 := len(queue) - 2
		for p1 >= 1 {
			//每次检查连个指针指向的数字
			if queue[p2] > queue[p1] {
				//如果左指针指向的数大于右边的，就交换
				fmt.Println(queue[p1], queue[p2])
				queue[p1], queue[p2] = change(queue[p1], queue[p2])
				fmt.Println(queue[p1], queue[p2])
			} else {
				//否则不做任何操作
				fmt.Println("ok sort:", queue[p1], queue[p2])
			}
			//两个指针继续向左移动一格
			p1 = p1 - 1
			p2 = p2 - 1
			fmt.Println("p1,p2 = ", p1, p2)
		}
		//将第一个元素是当前序列最小的数字，追加到返回序列。
		res = append(res, queue[0])
		queue = queue[1:]
	}
	return res
}

func change(i int, i2 int) (int, int) {
	return i2, i

}
