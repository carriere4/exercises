package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Println("Please enter 2 numbers: ")
	fmt.Scan(&num1, &num2)
	if num1 <= num2 {
		fmt.Println("The remainder is: ", num2%num1)
	} else {
		fmt.Println("The remainder is: ", num1%num2)
	}
}
