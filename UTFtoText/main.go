package main

//the object of this program is to take a string of utf8 codes and
//convert them into the rune represented by the codes
//only useful for roman alphabet as it assumes that every slice that starts with a 1
//is a 3 character slice

//Your first print is of string(65), where 65 is the codepoint for "A" so you get the string A.
//Your next prints are of substrings of your long string, for example str[0:2] which is "21",
//which is already a string and thus printed as 21.

//If you want to print the codepoint corresponding to 21
//you need to parse the string "21" as an integer and then convert that to a string.

import (
	"fmt"
	"strconv"
)

func main() {
	str := "8511033781167310910568528451100"
	//fmt.Printf("%v \n", string(65))
	//fmt.Println(len(str))
	for j, k := 0, 2; j < len(str)-1 && k < len(str); {
		if str[j:j+1] == "1" {
			if q, err := strconv.ParseInt(str[j:k+1], 10, 0); err == nil {
				fmt.Printf("%v", string(q))
				k = k + 3
				j = j + 3
			} //	fmt.Println(str[k:j])
		} else {
			if z, err := strconv.ParseInt(str[j:k], 10, 0); err == nil {
				fmt.Printf("%v", string(z))

				k = k + 2
				j = j + 2

			}
		}
	}
}

/* The code produces a "slice bounds out of range" panic. (https://play.golang.org/p/r-uswaokMx)
The problem here: for range moves through the string byte by byte, whereas k and j increase in steps of 2 or 3, respectively.
At some point, either k, j, or j+1 is larger than len(str),
which is when one of the slice operations on str fails.

Since your code moves through the string by increasing k and j,
I would suggest replacing the range operator with a loop with a condition,
like for k,j := 0,2; k < len(str) && j < len(str)-1;,
to ensure the loop stops before the slice operations go past the end of str.

*/
//if q, err := strconv.ParseInt(str[k:j], 10, 0); err == nil {
//  fmt.Printf("%v \n", string(str[q]))

//if z, err := strconv.ParseInt(str[k:j+1], 10, 0); err == nil {
//fmt.Printf("%v \n", string(z))
/*package main

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
}*/
