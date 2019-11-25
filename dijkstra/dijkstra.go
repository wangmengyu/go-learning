package main

import (
	"fmt"
	"math"
)

func main() {

	//存储有向图的三维MAP
	graph := make(map[string]map[string]float64)
	graph["start"] = make(map[string]float64)
	graph["start"]["a"] = 6
	graph["start"]["b"] = 2
	graph["a"] = make(map[string]float64)
	graph["a"]["fin"] = 1
	graph["b"] = make(map[string]float64)
	graph["b"]["a"] = 3
	graph["b"]["fin"] = 5

	//存储消耗的二维MAP
	costs := make(map[string]float64)
	costs["a"] = 6
	costs["b"] = 2
	costs["fin"] = math.Inf(1)

	//父节点二位MAP
	parent := make(map[string]string)
	parent["a"] = "start"
	parent["b"] = "start"
	parent["fin"] = ""

	//已经处理过的数据
	var processed = make(map[string]int)

	node := findLowestCostNode(costs, processed)
	fmt.Println(node)

	for node != "" {
		//取出当前最小消耗节点的消耗值
		cost := costs[node]
		fmt.Println("cost=", cost)
		//获取相邻节点
		neighbors := graph[node]
		for key, v := range neighbors {
			fmt.Printf("neighbour key=%s,val=%f", key, v)
			var newCost = cost + v
			if newCost < costs[key] {
				costs[key] = newCost
				parent[key] = node
			}
		}
		processed[node] = 1
		node = findLowestCostNode(costs, processed)
	}

	fmt.Println(costs, parent)

	//print the all path
	actKey := "start"
	var path []map[string]float64
	for len(parent) > 0 {
		for k, v := range parent {
			if v == actKey {
				p := make(map[string]float64)
				p[v] = costs[k]
				path = append(path, p)
				actKey = k
				delete(parent, k)
			}
		}
	}

	fmt.Println(path)

}

/**
找出当前消耗中最小消耗的节点
*/
func findLowestCostNode(costs map[string]float64, processed map[string]int) string {
	lowest := math.Inf(1)
	var lowestKey string
	for k, v := range costs {
		if v < lowest && processed[k] != 1 {
			lowest = v
			lowestKey = k
		}
	}
	return lowestKey
}
