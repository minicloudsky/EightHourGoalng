package main

// 声明一种行的数据类型，是int的一个别名
type myint int

type Book struct {
	title string
	auth string	
}
func changeBook(book Book) {
	book.auth = "666"
}

func changeBook2(book *Book) {
	book.auth = "666"
}

// func main() {
// 	var a myint = 10
// 	fmt.Println("a = ", a)
// 	fmt.Printf("type of a = %T\n", a)
// 	var book1 Book
// 	book1.auth = "Golang"
// 	book1.title = "zhang3"
// 	fmt.Printf("%v\n", book1)
// 	// 传值
// 	changeBook(book1) 
// 	fmt.Printf("%v\n", book1)
// 	// 修改结构体值，传指针
// 	changeBook2(&book1) 
// 	fmt.Printf("%v\n", book1)

// }