package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	go rLock()
	go wLock()
	go lock()
	time.Sleep(5 * time.Second)
}

func lock() {
	lock := sync.Mutex{}
	for i := 0; i < 3; i++ {
		lock.Lock()
		//defer lock.Unlock() //defer 最后执行解锁, 锁互斥, 代码报错
		fmt.Println("lock:", i)
		lock.Unlock()
	}
}

func rLock() {
	lock := sync.RWMutex{}
	for i := 0; i < 3; i++ {
		lock.RLock()
		defer lock.RUnlock() //读锁不互斥, 不建议在 defer 中写解锁
		fmt.Println("rLock:", i)
	}
}

func wLock() {
	lock := sync.RWMutex{}
	for i := 0; i < 3; i++ {
		lock.Lock()
		//defer lock.Unlock() //defer 最后执行解锁, 锁互斥, 代码报错
		fmt.Println("wLock:", i)
		lock.Unlock()
	}
}
