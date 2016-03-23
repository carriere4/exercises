package main

import "fmt"

func main() {
	var mult35 = 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			mult35 += i
		}

	}
	fmt.Println(mult35)
}
