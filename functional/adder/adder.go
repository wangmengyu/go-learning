package main

import "fmt"

/**
定义一个函数adder ,不需要参数， 返回一个 函数（这个函数将执行一些工作）
*/
func adder() func(int) int {
	s := 0
	return func(i int) int {
		s += i
		fmt.Println("s=", s)
		return s
	}
}

func main() {

	a := adder() // a 是一个函数，这个函数里有自己的初始化变量s=0。 返回的是一个函数。返回函数的功能是将s的值不断累加并且返回s的值
	for i := 1; i < 10; i++ {
		fmt.Println(a(i))
	}

}
