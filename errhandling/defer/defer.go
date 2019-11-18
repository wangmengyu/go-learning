package main

import (
	"bufio"
	"fmt"
	"github.com/wangmengyu/go-learning/functional/fib"
	"os"
)

func tryDefer() {
	defer fmt.Println(1) // defer 会在函数retrun 之前执行 defer 意思是推迟
	defer fmt.Println(2) // defer 是一个stack, 先进后出，这里会先输出2 再输出1
	fmt.Println(3)
	panic("err occurred") // 哪怕程序中遇到panic,还是会执行defer指定的语句

}

/**
写入文件
将FIB数序列写入到文件
*/
func writeFile(filename string) {
	//创建文件
	file, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	defer file.Close() // 最终会关闭文件

	write := bufio.NewWriter(file) //获得该文件的write对象
	defer write.Flush()            //最终会刷新write 对象

	f := fib.Fibonacci()
	for i := 0; i < 10; i++ {
		_, _ = fmt.Fprintln(write, f())
	}

}

func writeFileErr(filename string) {
	//创建文件
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)

	if err != nil {
		//fmt.Println("Err:",err.Error()) // 输出错误信息
		if pathError, ok := err.(*os.PathError); !ok {
			panic(pathError)
		} else {
			fmt.Printf("%s,%s,%s", pathError.Op, pathError.Path, pathError.Err)

		}
		return
	}

	defer file.Close() // 最终会关闭文件

	write := bufio.NewWriter(file) //获得该文件的write对象
	defer write.Flush()            //最终会刷新write 对象

	f := fib.Fibonacci()
	for i := 0; i < 10; i++ {
		_, _ = fmt.Fprintln(write, f())
	}

}

func main() {
	writeFile("fib.txt")
	writeFileErr("fib.txt")
}
