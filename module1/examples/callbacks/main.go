package main

import "fmt"

func main() {
	DoOperation(1, decrease)
	DoOperation(1, increase)
}

func increase(a, b int) {
	fmt.Println("increase result is:", a+b)
}

func DoOperation(y int, f func(int, int)) {
	f(y, 1)
}

func decrease(a, b int) {
	fmt.Println("decrease result is:", a-b)
}
