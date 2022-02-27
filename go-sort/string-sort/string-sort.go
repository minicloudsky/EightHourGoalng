package main

import (
	"fmt"
	"sort"
)

func main() {
	cities := []string{"shanghai", "shenzhen", "beijing",
		"Guangzhou", "Suzhou", "wuxi", "changzhou", "hangzhou"}
	fmt.Println("before sort: ", cities)
	sort.Strings(cities)
	fmt.Println("After sort: ", cities)
}
