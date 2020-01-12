package main

import (
	f "fmt"
)

func main() {

	var a, b, c int

	f.Scanf("%d %d", &a, &b)
	c = a * b

	if c%2 == 0 {
		f.Printf("Even\n")
	} else {
		f.Printf("Odd\n")
	}

}
