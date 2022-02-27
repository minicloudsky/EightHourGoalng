package main

import (
	"fmt"
	"sort"
)

func main() {
	balances := []float64{12.34, 89.1, 90.21, 21.12, 67.89, -90.1, 111, 21}
	fmt.Println("before sort: ", balances)
	sort.Float64s(balances)
	fmt.Println("after sort: ", balances)
}
