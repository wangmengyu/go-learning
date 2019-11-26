package main

import "fmt"

type weightPrice struct {
	weight int
	price  int
}

func main() {
	var weightArr = []int{1, 2, 3, 4}

	var products = make(map[string]weightPrice)
	products["g"] = weightPrice{price: 1500, weight: 1}
	products["m"] = weightPrice{price: 3000, weight: 4}
	products["n"] = weightPrice{price: 2000, weight: 3}

	//start to cal highest price for each cell
	//save every cell values
	table := make([][]int, len(products))

	i := 0
	for k, _ := range products {
		table[i] = make([]int, len(weightArr))
		j := 0
		for w := range weightArr {
			//last line max val
			// if has last line
			lastLineMax := 0
			if i > 0 {
				lastLineMax = table[i-1][j]
			} else {
				lastLineMax = 0
			}

			//check if weight is enough
			curLeftMaxPrice := 0
			if (products[k].weight) > w {
				//no need to get
				curLeftMaxPrice = 0
			} else {
				fmt.Println(k, i, j, w-products[k].weight)
				lastLineLeftMaxPrice := 0
				if w-products[k].weight >= 1 {
					lastLineLeftMaxPrice = table[i][w-products[k].weight]
				}
				curLeftMaxPrice = products[k].price + lastLineLeftMaxPrice
			}

			maxPrice := 0
			if curLeftMaxPrice < lastLineMax {
				maxPrice = lastLineMax
			} else {
				maxPrice = curLeftMaxPrice
			}
			table[i][j] = maxPrice
			j++
		}
		i++
	}

	fmt.Println(table)

}
