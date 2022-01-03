package main

import (
	"context"
	"fmt"
	"geekbang_cloud_native/module2/homework/httpserver"
	"sync"
	"time"
)

/*
课后练习 2.1
将练习 1.2 中的生产者消费者模型修改成为多个生产者和多个消费者模式
*/

func ProducerAndConsumer() {
	q := make(chan int, 10)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	lock := new(sync.Mutex)
	pwg := new(sync.WaitGroup)
	cwg := new(sync.WaitGroup)

	num := 0
	for i := 0; i < 5; i++ {
		pwg.Add(1)
		cwg.Add(1)
		go producer(ctx, q, &num, lock, pwg)
		go consumer(q, cwg)
	}

	pwg.Wait()
	close(q)
	cwg.Wait()
}

func producer(ctx context.Context, ch chan<- int, num *int, lock *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	t := time.NewTicker(500 * time.Millisecond)
	for range t.C {
		select {
		case <-ctx.Done():
			fmt.Println("producer work finish")
			return
		default:
			ch <- *num
			lock.Lock()
			*num++
			lock.Unlock()
		}
	}
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	t := time.Tick(time.Second)

	for num := range ch {
		<-t
		fmt.Println(num)
	}
	fmt.Println("consumer work finish")
}

func main() {
	ProducerAndConsumer()
	httpserver.ServerMain()
}
