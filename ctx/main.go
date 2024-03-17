package main

import (
	"context"
	"fmt"
	"time"
)

func RunTask(ctx context.Context) bool {
	childCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	time.Sleep(1 * time.Second)
	select {
	case <-childCtx.Done():
		// server closed, return
		return false
	default:
	}
	return true
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	result := RunTask(ctx)
	fmt.Println(result)
}
