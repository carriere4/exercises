package main

import "fmt"

func main() {
	fmt.Println("What is your name?")
	var yourname string
	fmt.Scan(&yourname)
	fmt.Printf("%v %v \n", "Hello,", yourname)
}
