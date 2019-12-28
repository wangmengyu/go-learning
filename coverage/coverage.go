package main

import (
	"fmt"
)

func main() {
	//需要覆盖的州
	statesNeeded := []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"}

	//各个站点可覆盖的州
	stations := make(map[string][]string)
	stations["kone"] = []string{"id", "nv", "ut"}
	stations["ktwo"] = []string{"wa", "id", "mt"}
	stations["kthree"] = []string{"or", "nv", "ca"}
	stations["kfour"] = []string{"nv", "ut"}
	stations["kfive"] = []string{"ca", "az"}

	//最终收集各个站点
	finalStations := make([]string, 0)

	//最优站点：覆盖最多未覆盖的站点
	for len(statesNeeded) > 0 && len(stations) > 0 {
		fmt.Println(len(statesNeeded))
		bestStation := ""
		statesCovered := make([]string, 0) //最优站点覆盖的未覆盖的州
		//对所有站台进行遍历
		for station, statesForStation := range stations {
			covered := intersection(statesForStation, statesNeeded) //获得当前站点具有的需求州列表
			if len(covered) > len(statesCovered) {
				bestStation = station
				statesCovered = covered
			}
		}
		//从需要的站台中减去已经覆盖的
		statesNeeded = subCovered(statesNeeded, statesCovered)
		finalStations = append(finalStations, bestStation)
		delete(stations, bestStation)
	}

	fmt.Println(finalStations)
}

func subCovered(needed []string, covered []string) []string {
	res := make([]string, 0)
	for _, n := range needed {
		found := false
		for _, c := range covered {
			if n == c {
				found = true
				break
			}
		}
		if found == false {
			res = append(res, n)
		}
	}
	return res
}

func intersection(station []string, needed []string) []string {
	res := make([]string, 0)
	for _, n := range needed {
		for _, s := range station {
			if n == s {
				res = append(res, n)
			}
		}
	}
	return res
}
