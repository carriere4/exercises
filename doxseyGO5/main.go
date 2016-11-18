package main

import "fmt"

/*func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("Panic")
}
*/

/*func zero(xPtr *int) {
	*xPtr = 0
}
func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // x is still 5
}
*/

/*func one(xPtr *int) {
	*xPtr = 1
}
func main() {
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // x is 1
}
*/

/*func half(x int) (int, bool) {
	if x%2 == 0 {
		return x / 2, true
	} else {
		return x / 2, false
	}
}

func main() {
	fmt.Println(half(1))
}
*/
//var max int

/*
func great(xs ...int) int {
	var max int
	for i, v := range xs {
		if i == 0 || v > max {
			max = v

		}
	}
	return max

}
func main() {
	fmt.Println(great(852, 4, 5, 101, 12, 78))

}
*/

/*
func fibonacci(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	fmt.Println(fibonacci(9))
}
*/

func f() {
	defer func() {

		if x := recover(); x != nil {

			fmt.Println("RECOVERED:", x)
		}
	}()
	panic("panic!")
}
func main() {
	//var y, x new(int)
	x := new(int)
	*x = 5
	fmt.Println(*x)
}
