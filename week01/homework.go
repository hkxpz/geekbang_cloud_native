package week01

import (
	"context"
	"fmt"
	"time"
)

/*
课后练习 1.1
编写一个小程序：
给定一个字符串数组
[“I”,“am”,“stupid”,“and”,“weak”]
用 for 循环遍历该数组并修改为
[“I”,“am”,“smart”,“and”,“strong”]
*/

func ChangeSlice() {
	me := []string{"I", "am", "stupid", "and", "weak"}

	for index, value := range me {
		switch value {
		case "stupid":
			me[index] = "smart"
		case "weak":
			me[index] = "strong"
		}
	}

	fmt.Println(me)
}

/*
课后练习 1.2
基于 Channel 编写一个简单的单线程生产者消费者模型：

队列：队列长度 10，队列元素类型为 int
生产者：每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
消费者：每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞
*/

func ProducerAndConsumer() {
	q := make(chan int, 10)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	go producer(ctx, q)
	go consumer(q)
	<-time.Tick(time.Second * 15)
}

func producer(ctx context.Context, ch chan<- int) {
	t := time.NewTicker(500 * time.Millisecond)
	num := 0
	for range t.C {
		select {
		case <-ctx.Done():
			fmt.Println("producer work finish")
			close(ch)
			return
		default:
			ch <- num
			num++
		}
	}
}

func consumer(ch <-chan int) {
	t := time.Tick(time.Second)
	for num := range ch {
		<-t
		fmt.Println(num)
	}
	fmt.Println("consumer work finish")
}
