package main

//proposed solution to https://projecteuler.net/problem=2
import "fmt"

var total int

func main() {
	for i, j := 0, 1; j < 4000000; i, j = i+j, i {
		//	fmt.Println(i)
		if i%2 == 0 {
			total = total + i
		}
	}
	fmt.Println(total)
}

// so I got completely stuck on how to initialize the fibonacci sequence
//so full disclosure: I got the fibonacci sequence in go from a web search
//and then I put the conditional statement in to get the answer

//
