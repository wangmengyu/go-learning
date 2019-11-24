package main

import "fmt"

func main() {
	graph := make(map[string][]string)
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "johnny"}

	var process = []string{"you"}
	people, ex := BFSearch(graph, process)
	if !ex {
		fmt.Println("not found people")
	} else {
		fmt.Println("find :", people)
	}

}

func BFSearch(graph map[string][]string, process []string) (string, bool) {
	for len(process) > 0 {
		people := process[0]
		process = process[1:]
		for _, v := range graph[people] {
			if endWithM(v) {
				fmt.Println("find people:", v)
				return v, true
			} else {
				process = append(process, v)
			}
		}
	}
	return "", false

}

func endWithM(s string) bool {
	lastChar := s[len(s)-1]
	return string(lastChar) == "m"

}
