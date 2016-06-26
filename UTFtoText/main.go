package main

//the object of this program is to take a string of utf8 codes and
//convert them into the rune represented by the codes
//only useful for roman alphabet as it assumes that every slice that starts with a 1
//is a 3 character slice
import "fmt"

var k = 0
var j = 2

func main() {
	str := "218526535870434111373812685"
	fmt.Printf("%v \n", string(65))
	//fmt.Println(len(str))
	for range str {
		if str[k:k+1] == "1" {
			fmt.Printf("%v \n", string(str[k:j+1]))
			k = k + 3
			j = j + 3
			//	fmt.Println(str[k:j])
		} else {
			fmt.Printf("%v \n", string(str[k:j]))
			k = k + 2
			j = j + 2

		}
	}
}
