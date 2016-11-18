package main

import "fmt"

var m []int
var mountainH int

func main() {
	for i := 0; i < 8; i++ {
		fmt.Scan(&mountainH)
		m = append(m, mountainH)
		for i, value := range m {
			var max int
			var imax int
			if m[i] > max {
				max = m[value]
				imax = i
				fmt.Println(imax)
			}

		}

		fmt.Println(m)
	}

}
