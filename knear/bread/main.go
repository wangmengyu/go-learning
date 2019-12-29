package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	data := make(map[string]map[string]int)

	//天气数据
	data["weather"] = make(map[string]int)
	data["weather"]["a"] = 5
	data["weather"]["b"] = 3
	data["weather"]["c"] = 1
	data["weather"]["d"] = 4
	data["weather"]["e"] = 4
	data["weather"]["f"] = 2
	data["weather"]["g"] = 4

	//是否周末
	data["is_weekend"] = make(map[string]int)
	data["is_weekend"]["a"] = 1
	data["is_weekend"]["b"] = 1
	data["is_weekend"]["c"] = 1
	data["is_weekend"]["d"] = 0
	data["is_weekend"]["e"] = 0
	data["is_weekend"]["f"] = 0
	data["is_weekend"]["g"] = 1

	//是否活动日
	data["is_act"] = make(map[string]int)
	data["is_act"]["a"] = 0
	data["is_act"]["b"] = 1
	data["is_act"]["c"] = 0
	data["is_act"]["d"] = 1
	data["is_act"]["e"] = 0
	data["is_act"]["f"] = 0
	data["is_act"]["g"] = 0
	disMap := make(map[string]float64)

	//已经知道的数量
	values := make(map[string]int)
	values["a"] = 300
	values["b"] = 225
	values["c"] = 75
	values["d"] = 200
	values["e"] = 150
	values["f"] = 50

	for name, _ := range data["weather"] {
		if name == "g" {
			continue
		}
		//计算name节点和g节点的距离的平方
		sum := float64(0)
		for _, item := range data {
			//当前遍历的指标，比较当前节点和目标节点的差的平方
			sum += math.Pow(math.Abs(float64(item[name])-float64(item["g"])), 2)
		}
		disMap[name] = sum
	}
	fmt.Println(disMap)

	//取出每个用例距离目标用例的距离值放入slice
	s := make([]float64, 0)
	for _, v := range disMap {
		s = append(s, v)
	}
	//对slice进行排序
	sort.Float64s(s)

	//存储结果的用例
	res := make(map[string]float64)
	resKeys := make([]string, 0)
	cnt := 0

	valSum := 0

	//存储最小的4个用例，请计算均值
I:
	for _, v1 := range s {
		for k, v2 := range disMap {
			if v1 != v2 {
				continue
			}
			//v1==v2时，查看是否已经收纳到结果map
			if _, ok := res[k]; !ok {
				res[k] = v1
				resKeys = append(resKeys, k)
				valSum += values[k]
				fmt.Println(k, v1)
				cnt++
				if cnt >= 4 {
					break I
				}
			}
		}
	}
	fmt.Println(resKeys)
	fmt.Printf("avg = %.2f ", float64(valSum/cnt))
}
