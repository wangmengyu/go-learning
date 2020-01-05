package main

import (
	"fmt"
	"strings"
)

type Subject struct {
	Title string //主题标题
	Trash int    //是否垃圾邮件
}

/**
以下主题给出了是否是垃圾邮件：
reset your password 0
you have won 1 million dollars 1
send me your password 1
peach prince sends you 10 million dollars 1
happy birthday 0
判断
collect your million dollars now 是垃圾邮件的概率是多少
*/

func main() {

	historyData := [...]Subject{
		{
			Title: "reset your password",
			Trash: 0,
		},
		{
			Title: "you have won 1 million dollars",
			Trash: 1,
		},
		{
			Title: "send me your password",
			Trash: 1,
		},
		{
			Title: "peach prince sends you 10 million dollars",
			Trash: 1,
		},
		{
			Title: "happy birthday",
			Trash: 0,
		},
	}

	//对每个单词统计它是垃圾的概率
	kwTrashCnt := make(map[string]int, 0)
	kwTrashTotal := 0
	wordCnt := make(map[string]int, 0)
	wordTotal := 0
	TrashCnt := 0

	//遍历样本
	for _, item := range historyData {
		//垃圾邮件的次数统计
		if item.Trash == 1 {
			TrashCnt++
		}
		//遍历每个单词
		words := strings.Split(item.Title, " ")
		for _, word := range words {
			//关键词包含在垃圾邮件中的次数统计
			if item.Trash == 1 {
				_, ok := kwTrashCnt[word]
				kwTrashTotal += 1
				if ok == false {
					kwTrashCnt[word] = 1
				} else {
					kwTrashCnt[word] += 1
				}
			}
			//所有单词出现的次数统计
			_, ok2 := wordCnt[word]
			wordTotal += 1
			if ok2 == false {
				wordCnt[word] = 1
			} else {
				wordCnt[word] += 1
			}
		}
	}

	//每个单词的概率集合
	pTrashWord := make(map[string]float64, 0)
	pAll := make(map[string]float64, 0)
	pTrash := float64(TrashCnt) / float64(len(historyData))
	fmt.Println("pTrash:", pTrash)
	for word, _ := range kwTrashCnt {
		//是垃圾邮件的概率
		val := float64(kwTrashCnt[word]) / float64(kwTrashTotal)
		pTrashWord[word] = val
		//总共的概率
		val2 := float64(wordCnt[word]) / float64(wordTotal)
		pAll[word] = val2
	}

	checkTitle := "collect your million dollars now"

	ckWords := strings.Split(checkTitle, " ")

	m := float64(1) //分子
	d := float64(1) //分母

	//P（具备的单词列表|垃圾邮件）
	//   = （P(垃圾) * P(单词1|垃圾) *P（单词2|垃圾）...P(单词N|垃圾)）
	//      /  P(单词1) * P（单词2） ... P(单词N)
	for _, word := range ckWords {
		//存在统计的单词
		_, ok := kwTrashCnt[word]
		if ok == true {
			m = m * pTrashWord[word]
			d = d * pAll[word]
			fmt.Println(pTrashWord[word], pAll[word])
		}
	}
	m *= pTrash
	fmt.Printf("res:%.9f", m/d)
}
