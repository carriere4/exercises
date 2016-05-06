package main

import "fmt"

func main() {
	a := true && false
	b := false && true
	c := false && false
	d := a || b || !c
	fmt.Println(c)
	fmt.Println(d)
}
