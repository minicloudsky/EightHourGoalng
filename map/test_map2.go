package main

import "fmt"

func printMap(cityMap map[string]string) {
	for key,value := range cityMap {
		fmt.Println("key= ", key,"value = ",value)
	}
}
func ChangeValue(cityMap map[string]string, key,value string) {
	cityMap[key] = value	
}

func main() {
	cityMap := make(map[string]string)
	// 添加
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"
	// 遍历
	for key,value := range cityMap {
		fmt.Println("key= ", key,"value = ",value)
	}
	delete(cityMap, "China")
	fmt.Println("--------")
	cityMap["Japan"] = "横须贺"
	// 遍历
	for key,value := range cityMap {
		fmt.Println("key= ", key,"value = ",value)
	}
	fmt.Println("~~~~~~~~~~~~")
	printMap(cityMap)
	ChangeValue(cityMap, "shanghai", "minhang")
	printMap(cityMap)

}