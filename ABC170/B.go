package main

import(
	"fmt"
)

//庭の動物の総数は X匹で、それらの足の総数は Y本である

func main() {
	var h, f int
	fmt.Scan(&h, &f)

	/* kame := (f/2) - h
	turu := f - kame

	if (kame+turu == h) && ((kame*4 + turu*2) == f) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	} */

	for i:=0; i<=h; i++ {
		for j:=0; j<=h-i; j++ {
			if i*4 + j*2 == f {
				fmt.Println("Yes")
			}
		}
	}
}