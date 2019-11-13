package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes你好啊" //UTF-8
	fmt.Println(len(s))
	//fmt.Printf("%X", []byte(s))

	//转成字节打印内容：一个中文需要3字节存储
	//59 65 73 E4 BD A0 E5 A5 BD E5 95 8A EF BC 81
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	//直接循环， 每个ch是一个rune
	for i, ch := range s {
		fmt.Printf("%d=>%X  ", i, ch)
	}
	fmt.Println()

	//UTF-8字数正确数值，6
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	//utf8.DecodeRune() 用法，将遇到的第一个字节转UTF-8字符并且返回长度
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c %d", ch, size)
	}

	fmt.Println()

	//用rune类型切s,即可获得正确的下标位置，和对应的中文字符
	for i, ch := range []rune(s) {
		fmt.Printf("(%d, %c)", i, ch)

	}

}
