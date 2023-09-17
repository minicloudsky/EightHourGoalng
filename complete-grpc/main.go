package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	//t, _ := time.Parse("2006-01-02 15:04:05", time.Now().String())
	//fmt.Printf("time: %v", t)
	fmt.Println(time.ParseInLocation("2006-01-02 15:04:05",
		time.Now().Format("2006-01-02 15:04:05"), time.Local))
}
