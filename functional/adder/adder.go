package adder

import "fmt"

/**
定义一个函数adder ,不需要参数， 返回一个 函数（这个函数将执行一些工作）
*/
func Adder() func(int) int {
	s := 0
	return func(i int) int {
		s += i
		fmt.Println("s=", s)
		return s
	}
}
