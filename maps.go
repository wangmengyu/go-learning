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

	//访问MAP的节点元素值，如果不存在KEY返回空数据
	fmt.Println("Getting values")
	courseName := m["course"]
	fmt.Println(courseName)
	courseName2 := m["courses"] //当取得MAP中不存在的KEY时，会返回空值，不报错
	fmt.Println(courseName2)

	//判断KEY是否存在
	courseName3, ok3 := m["course"]
	fmt.Println(courseName3, ok3)
	courseName4, ok4 := m["courses"]
	fmt.Println(courseName4, ok4)

	//通常判断KEY是否存在的写法
	key := "course"
	if courseName5, ok5 := m[key]; ok5 {
		fmt.Println(courseName5)
	} else {
		fmt.Printf("key does not exists:%s", key)
	}

	//删除元素
	fmt.Println("Delete values")
	name, nameOk := m["name"]
	fmt.Println(name, nameOk)
	delete(m, "name")
	name, nameOk = m["name"]
	fmt.Println(name, nameOk)

}
