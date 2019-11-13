package basic

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
  整数转二进制字符串
*/
func convertToBin(a int) string {
	res := ""
	for i := a; i != 0; i = i / 2 {
		mod := strconv.Itoa(i % 2) // strconv.Itoa 是整数转字符串
		fmt.Printf("m = %s \n", mod)
		res = mod + res
	}
	return res

}

/**
  逐行读取文件
  for 取代 while 用法，go 没有 while 语法
*/
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scaner := bufio.NewScanner(file)
	for scaner.Scan() { // 只保留退出条件，连段分号都可以省略，等效于while
		fmt.Println(scaner.Text())
	}
}

/**
  for语法也可以省略退出条件，意思是死循环。经常用于并发
*/
func forever() {
	for {
		fmt.Println("123")
	}
}

func main() {
	fmt.Println(convertToBin(5))  // 101
	fmt.Println(convertToBin(13)) // 13/2 = 6 余下1 6/2 = 3 余0 3/2 = 1 余下 1 , 最后是个1   1101 = 1+4+8
	//fmt.Println(13 / 2)
	printFile("abc.txt")
	//forever()

}
