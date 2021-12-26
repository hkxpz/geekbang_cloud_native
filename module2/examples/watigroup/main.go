package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var now time.Time
	now = time.Now()

	waitBySleep()
	fmt.Printf("waitBySleep exec time: %f s\n", time.Now().Sub(now).Seconds())
	now = time.Now()

	waitByChannel()
	fmt.Printf("waitByChannel exec time: %f s\n", time.Now().Sub(now).Seconds())
	now = time.Now()

	waitByWG()
	fmt.Printf("waitByWG exec time: %f s\n", time.Now().Sub(now).Seconds())
}

func waitBySleep() {
	for i := 0; i < 100; i++ {
		go fmt.Println(i)
	}
	time.Sleep(time.Second)
}

func waitByChannel() {
	ch := make(chan bool, 100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			ch <- true
		}(i)
	}

	for i := 0; i < 100; i++ {
		<-ch
	}
}

func waitByWG() {
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Println(i)
		}(i, &wg)
	}
	wg.Wait()
}
