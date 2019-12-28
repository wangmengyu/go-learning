package main

import "fmt"

func main() {
	//products 商品列表
	products := []string{"g", "n", "s"}
	//重量列表
	weights := make(map[string]int)
	weights["g"] = 1
	weights["n"] = 3
	weights["s"] = 4

	//价格列表
	prices := make(map[string]int)
	prices["g"] = 1500
	prices["n"] = 2000
	prices["s"] = 3000

	maxWeight := 4 // 最大负重
	cells := make([][]int, len(products))

	includeProduct := make(map[string]int, 0) //需要盗取的商品列表

	for i, product := range products {
		cells[i] = make([]int, maxWeight)
		for j := 0; j < maxWeight; j++ {
			//上一个单元格的值(cells[i-1]j)
			excludePrice := 0
			if i-1 >= 0 {
				excludePrice = cells[i-1][j]
			}
			//当前商品价格+剩余空间价值
			includePrice := 0
			leftPrice := 0
			if i-1 >= 0 && j-weights[product] >= 0 {
				leftPrice = cells[i-1][j-weights[product]]
			}
			includePrice = prices[product] + leftPrice

			if includePrice > excludePrice {
				//需要包含当前物品
				cells[i][j] = includePrice
				includeProduct[product] = 1
			} else {
				//不需要包含当前物品
				cells[i][j] = excludePrice
				delete(includeProduct, product)
			}

		}
	}

	maxVal := cells[len(products)-1][maxWeight-1]
	fmt.Println("max value =", maxVal)

	fmt.Println("need follow product list:")
	i := 1
	for product, _ := range includeProduct {
		fmt.Printf("%d.%s $%d ￡%d\n", i, product, prices[product], weights[product])
		i++
	}

}
