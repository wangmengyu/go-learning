package main

import (
	"fmt"
	"io/ioutil"
)

//swtch case 语法, 不需要在switch关键词后写表达式，也不要再每个case 内写break.
func grade(a int) string {
	g := ""
	switch {
	case a < 0 || a > 100:
		panic(fmt.Sprintf("Wrong score: %d", a)) // panic  一个报错语句，会中断程序的执行
	case a >= 90:
		g = "A"
	case a >= 80:
		g = "B"
	case a >= 70:
		g = "C"
	case a >= 60:
		g = "D"
	default:
		g = "E"
	}
	return g
}

func main() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil { // 判断条件不用小括号包住
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	//类for的if语法
	const filename2 = "bcd.txt"
	if contents2, err2 := ioutil.ReadFile(filename2); err != nil {
		fmt.Println(err2)
	} else {
		fmt.Printf("%s\n", contents2)
	}
	fmt.Println(grade(100))
	fmt.Println(grade(90))
	fmt.Println(grade(70))
	fmt.Println(grade(59))
	//fmt.Println(grade(-1))
	fmt.Println(grade(124))

}
