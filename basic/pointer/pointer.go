package main

import "fmt"

func swap(a,b *int) {
	var tmp int
	tmp = *a
	*a = *b
	*b = tmp
}
func main() {
	var a int =10
	var b int = 20
	fmt.Println("before swap, a= ",a," b= ",b)
	swap(&a,&b)
	fmt.Println("after swap, a= ",a," b= ",b)
	fmt.Println("a:",a, "&a: ",&a,"b: ",b,"&b: ",&b)
	var p *int
	p = &a
	fmt.Println("&a: ",&a)
	fmt.Println("p: ",p)
	var pp **int
	pp = &p
	fmt.Println("pp: ",pp)
	fmt.Println("*pp: ", *pp)
	fmt.Println("**pp: ", **pp)
}