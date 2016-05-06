package main

import "fmt"

func main() {
	n := greatnum(-128, -653.2, -43, -56, -87, -12, -45, -57)
	fmt.Println(n)
}

func greatnum(gn ...float64) float64 {
	fmt.Println(gn)
	var yuge float64
	for _, v := range gn {
		if v > yuge {
			yuge = v
		}
	}
	return yuge
}
