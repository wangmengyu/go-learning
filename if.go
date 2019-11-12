package main

import (
	"fmt"
	"io/ioutil"
)

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

}
