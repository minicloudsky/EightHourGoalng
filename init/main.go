package main

import (
	"GolangStudy/init/lib1"
	// . "GolangStudy/init/lib2" // 直接使用包内方法
	myLib2 "GolangStudy/init/lib2" // 给包起别名
)
func main(){
	lib1.Lib1Test()
	myLib2.Lib2Test()
}