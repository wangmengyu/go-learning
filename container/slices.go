package container

import "fmt"

//修改切片的第一个元素为，
func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println("arr[2:6]=", arr[2:6])
	fmt.Println("arr[:6]=", arr[:6])
	s1 := arr[2:]
	fmt.Println("s1=", s1)
	s2 := arr[:]
	fmt.Println("s2=", s2)

	updateSlice(s1)
	fmt.Println("[after update s1]", s1)
	fmt.Println("[origin s]", arr)

	updateSlice(s2)
	fmt.Println("[after update s2]", s2)
	fmt.Println("[origin s]", arr)

	fmt.Println("Reslice")
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("Extending slice")
	arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr)
	s1 = arr[2:6]
	fmt.Println("s1:", s1)
	s2 = s1[3:5]
	fmt.Println("s2:", s2) // [5,6] 切片会尝试追踪到原数组，查看是否有值，即使中间的视图已越界

	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", //cap() 函数 表示容量
		s1, len(s1), cap(s1))

	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n",
		s2, len(s2), cap(s2))

	//append, 对切片的append，会修改到原数组上, 如果超过原来数组上界，会新分配内存，但是不修改原数组长度值
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3,s4,s5 = ", s3, s4, s5)
	fmt.Println("arr", arr)

}
