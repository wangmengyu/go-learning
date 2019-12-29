package main

import (
	"fmt"
	"math"
)

/**
  K邻近算法
*/

func main() {
	data := make(map[string]map[string]int)
	data["comedy"] = make(map[string]int)
	data["comedy"]["p"] = 3
	data["comedy"]["j"] = 4
	data["comedy"]["m"] = 2
	data["act"] = make(map[string]int)
	data["act"]["p"] = 4
	data["act"]["j"] = 3
	data["act"]["m"] = 5
	data["life"] = make(map[string]int)
	data["life"]["p"] = 4
	data["life"]["j"] = 5
	data["life"]["m"] = 1
	data["dracula"] = make(map[string]int)
	data["dracula"]["p"] = 1
	data["dracula"]["j"] = 1
	data["dracula"]["m"] = 3
	data["love"] = make(map[string]int)
	data["love"]["p"] = 4
	data["love"]["j"] = 5
	data["love"]["m"] = 1

	//计算p和j的距离
	//五个维度
	var sum float64
	for filmType, _ := range data {
		sum += math.Pow(math.Abs(float64(data[filmType]["p"]-data[filmType]["j"])), 2)
	}
	dis := math.Sqrt(sum)
	fmt.Println("distance of p & j = ", dis)

}
