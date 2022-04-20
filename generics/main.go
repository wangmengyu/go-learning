package main

import "fmt"

func main() {
	ints := map[string]int{
		"1st": 1,
		"2nd": 2,
	}

	floats := map[string]float64{
		"1st": 1.1,
		"2nd": 2.2,
	}

	res := SumIntOrFloat(ints)
	fmt.Println(res)

	res2 := SumIntOrFloat(floats)
	fmt.Println(res2)

}

type Number interface {
	int | float64
}

func SumIntOrFloat[K comparable, V Number](m map[K]V) V {
	var s V
	for k, v := range m {
		s += v
		fmt.Println(k)
	}
	return s
}
