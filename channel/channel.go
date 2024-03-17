package main

import (
	"fmt"
	"sort"
)

func main() {
	slices := []int{13, 354, 454, 56, 1, 24, 90}
	sort.Ints(slices)
	fmt.Println(slices)
}
