package main

import "fmt"

var imax int
var max int
var mounts = []int{212, 13, 11, 24, 10, 18, 18, 20}

//var mountainH int

func main() {
	for i := 0; i < 8; i++ {
		//fmt.Scan(&mountainH)
		//m = append(m, mountainH)

		for j, mount := range mounts {
			//var max int
			//var imax int
			if mount > max {
				max = mount
				imax = j
				//fmt.Println(imax)
			}

		}

		//fmt.Println(mounts)
	}
	fmt.Println(imax)
	//fmt.Println(max)
	//fmt.Println(mounts)
}
