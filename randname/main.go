package main

import (
	"fmt"
	"github.com/tjarratt/babble"
)

func main() {
	babbler := babble.NewBabbler()
	babbler.Count = 1
	fmt.Println(babbler.Babble())
}
