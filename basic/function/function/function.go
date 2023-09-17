package function

// Add 具名函数
func Add(a, b int) int {
	return a + b
}

// Add2 匿名函数
var Add2 = func(a, b int) int {
	return a + b
}

// Swap 多个参数和多个返回值
func Swap(a, b int) (int, int) {
	return b, a
}

// Sum 可变数量的参数
// more 对应 []int 切片类型
func Sum(a int, more ...int) int {
	for _, v := range more {
		a += v
	}
	return a
}

// Find 函数返回值也可以命名
func Find(m map[int]int, key int) (value int, ok bool) {
	value, ok = m[key]
	return
}

// Inc 返回值命名了，可以通过名字来修改返回值，也可以通过defer语句在return语句之后修改返回值
func Inc() (v int) {
	// defer语句延迟执行了一个匿名函数，因为这个匿名函数捕获了外部函数的局部变量v，这种函数我们一般叫闭包。
	// 闭包对捕获的外部变量并不是传值方式访问，而是以引用的方式访问
	defer func() { v++ }()
	return 42
}
