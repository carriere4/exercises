package main

//this takes a string and returns the UTF8 numeric representations
import (
	"fmt"
)

func main() {
	str := "UNintimidated"
	for _, r := range str {
		//c := string(r)
		fmt.Printf("%d", r)
	}
}
