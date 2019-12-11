package main

import (
	"fmt"
	"regexp"
)

const text = `
my email is mengyu.wang.0124@gmail.com
email1 is abc@def.org 
email2 is kkk@qq.com 
email3 is ddd@abc.com.cn
`

func main() {

	//创建正则匹配工具
	//要提取的部分，周围都加上小括号
	reg := regexp.MustCompile(`([a-zA-Z0-9.]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)

	//查找第一个匹配到的字符串
	s := reg.FindString(text)
	//查找所有的匹配结果
	arr := reg.FindAllString(text, -1)

	//查找所有小括号提取的部分
	subArr := reg.FindAllStringSubmatch(text, -1)

	fmt.Println(s)
	fmt.Println(arr)

	for _, match := range subArr {
		fmt.Println(match)
	}

}
