package main

// func main(){
// 	// tty: pair<type:*os.File, value:"/dev/tty" 文件描述符>
// 	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
// 	if err !=nil {
// 		fmt.Println("open file error!", err)
// 		return
// 	}
// 	fmt.Println(tty.Name())
// 	// r: pair<type: *os.File, value: "/dev/tty" 文件 描述符>
// 	var r io.Reader
// 	// r: pair<type:*os.File,value:"/dev/tty" 文件描述符>
// 	r = tty
// 	// w: pair<type:*os.File,value:"/dev/tty" 文件描述符>
// 	w := r.(io.Writer)
// 	w.Write([]byte("Hello! this is a test!\n"))
// }