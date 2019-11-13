package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d \n", s,
		len(s), cap(s))
}
func main() {
	fmt.Println("Creating slice")
	var s []int // 定义切片变量， zero value = nil
	//给切片追加100个元素
	for i := 0; i < 100; i++ {
		s = append(s, 2*i+1)
		//printSlice(s) // 每次扩充以2的指数增加，2，4，8，16
	}
	fmt.Println(s)

	//创建切片时指定元素值。
	s1 := []int{2, 4, 6, 8}
	printSlice(s1)
	//创建切片时候，指定长度
	s2 := make([]int, 16)
	printSlice(s2)
	//创建切片时候，指定长度与容量
	s3 := make([]int, 10, 32)
	printSlice(s3)

	//切片的复制，复制时，会保留目标切片已存在的元素，将新切片的内容覆盖掉目标切片的元素，
	fmt.Println("Copying slice")
	copy(s2, s1)
	fmt.Println("Copy s1 to s2: s2=", s2)

	//删除第3个位置的元素
	fmt.Println("Delete slice")
	s2 = append(s2[:3], s2[4:]...) // s2[4:]后追加3个点表示将4位置后面的元素全部展开为一个个参数传入到函数中
	printSlice(s2)

	//弹出第一个元素
	fmt.Println("Pop from front")
	front := s2[0]
	fmt.Println("front = ", front)
	s2 = s2[1:]
	printSlice(s2)

	//弹出最后一个元素
	fmt.Println("Pop from end")
	end := s2[len(s2)-1]
	fmt.Println("end = ", end)
	s2 = s2[:len(s2)-1]
	printSlice(s2)

}
