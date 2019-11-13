package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var ( //函数外部定义变量，作用于整个pakage内
	aa = 11
	bb = 22
	cc = "hi"
)

/**
不付初始值会自动有初始值，整数0，字符串空串
*/
func variableZeroValue() {
	var a int //
	var s string
	fmt.Println(a, s)
	fmt.Printf("%d %q", a, s)

}

func variableInitValue() {
	var a, b = 3, 4 //
	var s = "abc"
	fmt.Println(a, s, b)
}

/**
多个变量赋值，不用指定类型
*/
func variableTypeDeduction() {
	var a, b, c, d = 1, 2, true, "abc"
	fmt.Println(a, b, c, d)
}

/**
  定义变量缩写形式为 ：注意： 这种方式只能用于函数内部定义变量，外部不可以
    变量名1,变量名2,变量名3,... := 值1，值2，值3...
    不需要var 关键字， 推荐使用这种方式定义
*/
func variableShorter() {
	a, b, c, d := 1, 2, 3, "abc"
	fmt.Println(a, b, c, d)
}

//欧拉公式
func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))

	d := cmplx.Pow(math.E, 1i*math.Pi) + 1
	fmt.Println(d)

	e := cmplx.Exp(1i*math.Pi) + 1
	fmt.Printf("%.3f", e)

}

//类型必须强制转换
func triangle() {
	a, b := 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a) + float64(b*b)))
	fmt.Println(c)
}

/**
常量定义时候，不标记类型，可以是任何类型，不需要强制转换类型
*/
func consts() {
	const a, b = 3, 4
	const d = "abc"
	/**
	  也可以一组一组定义：
	  const (
	     a,b = 3,4
		 d = "abc"
	   )
	*/
	c := math.Sqrt(a*a + b*b)
	fmt.Println(c)
}

/**
枚举类型
*/
func enums() {
	const ( //  指定每个值
		cpp    = 1
		java   = 2
		python = 3
		golang = 4
	)

	const (
		php = iota // 自动从0每个元素值开始自增
		c
		js
		ruby
	)

	const (
		b = 1 << (10 * iota) // 左移运算符，原始值1二进制=0001. 左移10次后等于 1000000000=2的十次方，1024，以此类推
		kb
		mb
		gb
	)

	fmt.Println(cpp, java, python, golang)
	fmt.Println(php, c, js, ruby)
	fmt.Println(b, kb, mb, gb)

}
func main() {
	fmt.Println("hello123")
	//variableZeroValue()
	variableInitValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, cc)
	euler()
	triangle()
	consts()
	enums()

}
