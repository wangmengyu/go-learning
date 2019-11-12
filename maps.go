package main

import "fmt"

func main() {
	//定义一个KEY是String,Val也是String的map,定义时就初始化好值
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"quality": "notbad",
	}

	//用make函数定义一个空的map, key是string,val是int
	m2 := make(map[string]int)

	//用关键字var定义一个空的map, key是string,val是int
	var m3 map[string]int
	fmt.Println(m, m2, m3)

	//遍历map
	fmt.Println("Travers map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	//取得MAP的节点元素值，如果不存在KEY返回空数据
	fmt.Println("Getting values")
	courseName := m["course"]
	fmt.Println(courseName)

}
