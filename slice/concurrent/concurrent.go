package main

import (
	"fmt"
	"time"
)

var slices = []int{1, 2, 3, 4, 5, 6, 7, 8}

func add() {
	slices = append(slices, 9)
	fmt.Println("slices in add: \n", slices)
}

func modify() {
	slices[len(slices)-1] = 11
	fmt.Println("slices in modify: \n", slices)
}

func main() {
	go add()
	go modify()
	time.Sleep(1 * time.Second)
	fmt.Println("---main slices: ", slices)
}
