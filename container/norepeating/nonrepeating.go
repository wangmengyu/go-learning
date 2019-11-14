package main

import (
	"fmt"
	"reflect"
)

//寻找最长不含有重复字符的子串(第一次出现的最大长度)
/**
a,b,c,a,b,c,b,b -> abc
算法：
  设定一个MAP存储每个字符出现的最大位置
  从头遍历每个字符
    判断是否存在已收集的该字符的最大位置，如果存在，判断是否大于开始位置，
    当大于开始位置时候，更新start位置
    如果当前方位位置减去开始位置大于最大长度，则更新最大长度
*/
func lenNonRepSubStr(s string) (int, string, int) {

	lastOccurred := make(map[rune]int) //每个字母最后出现的位置KEY是字母，val是位置
	start := 0
	maxLen := 0
	runeSli := []rune(s)
	subStr := runeSli
	for i, ch := range runeSli {
		fmt.Println(reflect.TypeOf(ch), i, string(ch))
		lastI, ok := lastOccurred[ch]

		if ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLen {
			maxLen = i - start + 1
			subStr = runeSli[start : i+1]
			fmt.Println("[sub str]=", string(subStr))
			fmt.Println(start, i+1)
		}
		lastOccurred[ch] = i
	}
	return start, string(subStr), maxLen

}

func main() {
	start, substr, maxLen := lenNonRepSubStr("上海自来水来自海上")
	fmt.Println(start, substr, maxLen)

}
