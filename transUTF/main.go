package main

//this takes a string and returns the UTF8 numeric representations
import "fmt"

func main() {
	fmt.Print("Enter text without spaces: ")
	var input string
	fmt.Scanln(&input)

	for _, r := range input {
		//c := string(r)
		fmt.Printf("%d", r)
	}
	fmt.Println()
}
