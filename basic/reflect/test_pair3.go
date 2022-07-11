package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

// 具体类型
type  Book struct {

}

func (this *Book) ReadBook () {
	fmt.Println("Read a book!")
}

func (this *Book) WriteBook() {
	fmt.Println("Write a book!")
}

// func main() {
// 	// b: pair<type:Book,value:book{}地址>
// 	b := &Book{}
// 	var r Reader
// 	r = b
// 	r.ReadBook()

// 	var w Writer
// 	// r: pait<type:Book,value:Book{}地址>
// 	w = r.(Writer) // 此处断言为何会成功？因为 w r 具体的type是一致的
// 	fmt.Println(w)
// }