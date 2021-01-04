package main

import (
	"fmt"
	"math/rand"
)

/**
  协同过滤法的离线实验
*/
func main() {
	//先准备数据,
	data := InitData()
	train, test := SplitData(data, 8, 1, 42)
	fmt.Println(len(train))
	fmt.Println(len(test))
}

type DataItem struct {
	uid    int
	itemId int
	score  int
}

/**
60个用户, 40个电影
*/
func InitData() []DataItem {
	res := make([]DataItem, 0)
	for uid := 1; uid <= 60; uid++ {
		for itemId := 1; itemId <= 40; itemId++ {
			score := rand.Intn(5) + 1
			res = append(res, DataItem{
				uid:    uid,
				itemId: itemId,
				score:  score,
			})
		}
	}

	return res
}

/**
  从一堆数据中抽取一份样本, 其余的作为测试模型数据
*/
func SplitData(data []DataItem, M, k int, seed int) (map[int][]DataItem, map[int][]DataItem) {
	test := make(map[int][]DataItem)  // 测试数据集
	train := make(map[int][]DataItem) // 训练数据集
	rand.NewSource(int64(seed))
	for _, v := range data {
		if _, testOk := test[v.uid]; !testOk {
			test[v.uid] = make([]DataItem, 0)
		}
		if _, trainOk := train[v.uid]; !trainOk {
			train[v.uid] = make([]DataItem, 0)
		}
		if int(rand.Int31n(int32(M))) == k {
			test[v.uid] = append(test[v.uid], v)
		} else {
			train[v.uid] = append(train[v.uid], v)
		}
	}
	fmt.Println("[train]:", train)
	fmt.Println("[test]:", test)
	return train, test

}

/**
  召回率:
   秒数有多少比例的用户-物品评分包含在最终的推荐列表中
*/
func Recall() {

}

/**
  准确率:
  最终的推荐列表中有多少比例是发生过的用户-物品 评分记录
*/
func Precision() {

}

/**
覆盖率 :
  最终的推荐列表中包含多大比例的物品
  如果所有物品都被推荐给至少一个用户, 那么覆盖率是100%
*/
func Coverage() {

}
