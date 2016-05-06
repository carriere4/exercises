package main

import "fmt"

func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}

func main() {
	var h int
	var even bool
	h, even = half(5)
	fmt.Println(h, even)
}
