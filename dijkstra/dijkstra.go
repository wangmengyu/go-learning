package main

import (
	"fmt"
	"math"
)

func main() {
	//点到点的距离
	graph := make(map[string]map[string]int)
	graph["start"] = make(map[string]int)
	graph["start"]["a"] = 6
	graph["start"]["b"] = 2
	graph["a"] = make(map[string]int)
	graph["a"]["fin"] = 1
	graph["b"] = make(map[string]int)
	graph["b"]["a"] = 3
	graph["b"]["fin"] = 5
	graph["fin"] = make(map[string]int)

	//消耗表
	costs := make(map[string]int)
	costs["a"] = 6
	costs["b"] = 2
	costs["fin"] = math.MaxInt32

	//parent表
	parent := make(map[string]string)
	parent["a"] = "start"
	parent["b"] = "start"
	parent["fin"] = ""

	//processed表
	processed := make(map[string]int, 0)

	//寻找消耗最低的节点
	for {
		node, ok := getMinCost(costs, processed)
		if ok == false {
			//没有要处理的节点
			break
		}
		fmt.Println("min cost node:", node)
		//标记已经处理
		processed[node] = 1

		//访问该结点的邻居结点
		for neighbor, cost := range graph[node] {
			newCost := cost + costs[node] // 新距离 = 当前节点消耗+当前节点到邻节点的距离
			if newCost < costs[neighbor] {
				// 新距离比消耗表的小，需要更新 消耗表和parent表
				costs[neighbor] = newCost
				parent[neighbor] = node
			}
		}
	}

	fmt.Println("costs", costs)

	//打印最终路径(parent表),倒过来从，fin开始搜索，每次找它的父节点进行路径记录。直到start点

	next := "fin"
	path := make([]string, 0)
	path = append(path, next)

	for {
		p, ok := parent[next]
		if ok == false {
			break
		}
		currentPath := []string{p}
		path = append(currentPath, path[:]...)
		next = p
	}

	fmt.Println(path)
}

/**
  找出当前消耗最低的节点
*/
func getMinCost(costs map[string]int, processed map[string]int) (string, bool) {
	minVal := math.MaxInt32
	findNode := ""
	find := false
	for name, cost := range costs {
		if _, ok := processed[name]; ok {
			continue
		}
		if cost < minVal {
			minVal = cost
			findNode = name
			find = true
		}
	}
	return findNode, find
}
