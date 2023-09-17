package function

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	fmt.Println(Add(1, 1))
}

func TestSwap(t *testing.T) {
	fmt.Println(Swap(1, 10))
}

func TestSum(t *testing.T) {
	fmt.Println(Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

func TestChangableParam(t *testing.T) {
	var a = []interface{}{123, "abc"}

	fmt.Println(a...) // 123 abc 等价于直接调用Print(123, "abc")
	fmt.Println(a)    // [123 abc] 等价于直接调用Print([]interface{}{123, "abc"})

}

func TestInc(t *testing.T) {
	fmt.Println(Inc())
}
