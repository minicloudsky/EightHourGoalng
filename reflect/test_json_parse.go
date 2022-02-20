package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// type Movie struct {
// 	Title string  `json:"title"`
// 	Year int      `json:"year"`
// 	Price int `json:"rmb"`
// 	Actors []string `json:"actors"`
// }


type Movie struct {
	Title string  // 不用结构体标签时候，编解码出来的是大写的
	Year int     
	Price int `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 2000,10,[]string{"zhouxingchi","zhangbozhi"}}
	
	// 编码的过程，结构体 ---> json
	jsonStr ,err := json.Marshal(movie)
	if err!=nil {
		fmt.Println("json parse err!")
		return
	} else {
		fmt.Println("movie: ", jsonStr,"\nmovie type: ", reflect.TypeOf(jsonStr), "\njsonStr: ", string(jsonStr))
	}
	fmt.Printf("movie: %s\n", jsonStr)
	// json解码过程， jsonStr ---> 结构体
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err!=nil{
		fmt.Println("json unmarsharll err!", err)
		return 
	}
	fmt.Println("my_movie: ", myMovie)


}