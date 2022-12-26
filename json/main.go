package main

import (
	"encoding/json"
	"fmt"
)

type OrderItem struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Qty   int     `json:"qty"`
}
type Order struct {
	Id    string       `json:"id"`
	Item  []*OrderItem `json:"item"`
	Total float64      `json:"total"` // 总计
}

func main() {
	/*
		o := Order{
			Id: "1234",
			Item: []*OrderItem{
				{
					Name:  "xxx",
					Price: 10,
					Qty:   1,
				},
				{
					Name:  "nnn",
					Price: 10,
					Qty:   1,
				},
			},
			Total: 15,
		}
		fmt.Printf("%+v \n", o)
		b, err := json.Marshal(o)
		if err != nil {
			panic(err.Error())
		}

		fmt.Printf("%s\n", b)

		str := `{"id":"1234","item":[{"name":"xxx","price":10,"qty":1},{"name":"nnn","price":10,"qty":1}],"total":15}`
		var ord Order
		err = json.Unmarshal([]byte(str), &ord)
		if err != nil {
			panic(err.Error())
			return
		}
		fmt.Printf("%+v", ord)

	*/
	ParseNLP()

}

// ParseNLP 调用阿里自然语言处理 API
func ParseNLP() {
	str := `[{"id":"0","word":"阿里巴巴","tags":["机构-机构半称","产品-品牌"]},{"id":"1","word":"集团","tags":["基本词-中文","产品类型修饰词"]},{"id":"2","word":"的","tags":["基本词-中文","文体娱乐类-flash作品"]},{"id":"3","word":"使命","tags":["基本词-中文","文体娱乐类-书文课程类","软件-纯软件名","产品类型修饰词"]},{"id":"4","word":"是","tags":["基本词-中文","文体娱乐类-flash作品"]},{"id":"5","word":"让","tags":["基本词-中文","文体娱乐类-flash作品"]},{"id":"6","word":"天下","tags":["产品-品牌","基本词-中文","软件-纯软件名"]},{"id":"7","word":"没有","tags":["基本词-中文","产品类型修饰词"]},{"id":"8","word":"难做","tags":["人民日报词汇"]},{"id":"9","word":"的","tags":["基本词-中文","文体娱乐类-flash作品"]},{"id":"10","word":"生意","tags":["基本词-中文","网站-其他","文体娱乐类-书文课程类","文体娱乐类-报纸杂志类","文体娱乐类-戏剧歌曲类","机构-机构特指"]},{"id":"11","word":"。","tags":[]}]`
	data := make([]struct { // 直接定义一种结构体进行使用
		Id   string   `json:"id"`
		Word string   `json:"word"`
		Tags []string `json:"tags"`
	}, 0)
	err := json.Unmarshal([]byte(str), &data)
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
		return
	}

	fmt.Printf("%+v\n", data[1].Tags)

}
