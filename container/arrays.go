package container

import "fmt"

/**
  打印数组
*/
func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

/**
编辑数组的元素值
*/
func editArray(arr *[5]int) {
	arr[0] = 100
}

func main() {
	var arr1 [5]int                    // var 定义的不用设置初始值
	arr2 := [3]int{1, 3, 5}            // := 定义的必须设置初始值
	arr3 := [...]int{1, 20, 30, 40, 5} // ... 不用明确给出元素长度,让编译器自动计算
	fmt.Println(arr1, arr2, arr3)
	var grid [4][5]int // 4行5列的整数数组
	fmt.Println(grid)

	//遍历数组，range的用法
	for i, v := range arr3 {
		fmt.Println(i, v)
	}
	//如果不需要i可以用下划线代替忽略使用
	for _, v := range arr3 {
		fmt.Println(v)
	}

	//找出值最大的元素
	maxI := 0
	maxV := 0
	for i, v := range arr3 {
		if v > maxV {
			maxI, maxV = i, v
		}
	}
	fmt.Println(maxI, maxV)

	printArray(arr1)

	fmt.Println(arr1)

	editArray(&arr1)

	fmt.Println(arr1)

}
