package main

import (
	"fmt"
	"math"
	"sort"
)

/**
  K邻近算法
*/
type bread struct {
	weather   int
	isWeekend int
	isAct     int
	num       int
}

func main() {
	breadGroup := make(map[string]bread, 6)
	breadGroup["a"] = bread{weather: 5, isWeekend: 1, isAct: 0, num: 300}
	breadGroup["b"] = bread{weather: 3, isWeekend: 1, isAct: 1, num: 225}
	breadGroup["c"] = bread{weather: 1, isWeekend: 1, isAct: 0, num: 75}
	breadGroup["d"] = bread{weather: 4, isWeekend: 0, isAct: 1, num: 200}
	breadGroup["e"] = bread{weather: 4, isWeekend: 0, isAct: 0, num: 150}
	breadGroup["f"] = bread{weather: 2, isWeekend: 0, isAct: 0, num: 50}

	cmpBread := bread{weather: 4, isWeekend: 1, isAct: 0, num: 0}

	sortVals := make([]float64, len(breadGroup))

	fmt.Println(len(breadGroup))
	i := 0
	sortSeq := make(map[string]float64, len(breadGroup))
	for k, v := range breadGroup {
		//获得所有点与参考点的距离列表
		d := math.Sqrt(math.Pow(float64(v.weather-cmpBread.weather), 2) + math.Pow(float64(v.isWeekend-cmpBread.isWeekend), 2) + math.Pow(float64(v.isAct-cmpBread.isAct), 2))
		fmt.Println(d)
		sortSeq[k] = d
		sortVals[i] = d
		i++
	}

	//对距离进行排序
	//quick sort for seq
	sort.Float64s(sortVals)

	fmt.Println(sortVals)

	keys := make(map[string]float64, 4)

	cnt := 0
	var sum float64
	sum = 0

	//获得最少距离的4个案例
	for _, score := range sortVals {
		if cnt >= 4 {
			break
		}
		for k, v := range sortSeq {
			_, ok := keys[k]
			if v == score && !ok {
				keys[k] = v
				curBread := breadGroup[k]
				sum = float64(curBread.num) + sum
				cnt++
			}
		}
	}

	//计算平均值
	fmt.Println(sum / 4)

}
