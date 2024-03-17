package main

import (
	"errors"
	"fmt"
	"sync"
)

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *Stack[T]) Pop() (t T, err error) {
	if len(s.data) == 0 {
		return t, errors.New("stack is empty")
	}
	res := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return res, nil
}

func main() {
	var stack Stack[int]
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	item, err := stack.Pop()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Pop item:", item)
	}
	item, err = stack.Pop()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Pop item:", item)
	}
	rw := sync.RWMutex{}
	rw.RLock()
}
