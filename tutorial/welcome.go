package main

import (
	f "fmt"
)

func main() {

	var a, b, c int
	var s string

	f.Scanf("%d", &a)
	f.Scanf("%d %d", &b, &c)
	f.Scanf("%s", &s)

	f.Printf("%d %s\n", a+b+c, s)

}
