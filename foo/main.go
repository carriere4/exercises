package main

import "fmt"

func foo(number ...int) {
	fmt.Println(number)
}

func main() {
	foo(1, 2, 3)
	foo(1, 4, 5, 98845, 8574)
}
