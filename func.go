package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

/**
  多个参数
*/
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b) // 当第二个返回值不想用。，指想用第一个，可以用_给第二个变量返回，防止编译器报错
		return q, nil
	default:
		return 0, fmt.Errorf("no supported op : %s", op)
	}

}

/**
  已函数作为参数
*/
func apply(op func(int, int) int, a, b int) int {

	p := reflect.ValueOf(op).Pointer()    //获得op函数的指针
	opName := runtime.FuncForPC(p).Name() //获得运行中的函数方法名

	fmt.Printf("Calling func %s with args"+
		"(%d, %d)", opName, a, b)
	return op(a, b)

}

/**
多个返回值，一一对应，并可以设置返回值名用于编辑器接收, 但是一般不需要
*/
func div(a, b int) (int, int) {
	return a / b, a % b
}

/**
可变数量参数，类型一致。在参数后加...
*/
func sum(numbers ...int) int {
	s := 0
	for i := range numbers { //遍历每个参数
		s += numbers[i]
	}
	return s
}

func main() {
	r, e := eval(13, 4, "++")
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(r)
	}

	q, r := div(13, 4)
	fmt.Println(q, r)
	/**
	apply 方法接收的第一个参数是一个匿名函数
	*/
	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

	fmt.Println(sum(1, 2, 3))

}
