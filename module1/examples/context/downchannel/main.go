package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan int, 10)
	done := make(chan bool)
	defer close(messages)

	//consumer
	go func() {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			select {
			case <-done:
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Printf("send message: %d\n", <-messages)
			}
		}
	}()

	//producer
	for i := 0; i < 10; i++ {
		messages <- i
	}

	time.Sleep(5 * time.Second)
	close(done)
	time.Sleep(time.Second)
	fmt.Println("main process exit!")
}