package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	baseCtx := context.Background()
	ctx := context.WithValue(baseCtx, "a", "b")

	go func(ctx context.Context) {
		fmt.Println(ctx.Value("a"))
	}(ctx)

	timeOutCtx, cancel := context.WithTimeout(baseCtx, time.Second)
	defer cancel()
	go func(ctx context.Context) {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			select {
			case <-ctx.Done():
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Println("enter default")
			}
		}
	}(timeOutCtx)

	select {
	case <-timeOutCtx.Done():
		time.Sleep(time.Second * 2)
		fmt.Println("main process exit!")
	}
}
