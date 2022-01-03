package main

import "fmt"

func main() {
	f := func(x, y int) {
		fmt.Println(x + y)
	}
	f(1, 2)
	a(1, f)

}

func a(x int, f func(x, y int)) {
	f(x, 6)
}
