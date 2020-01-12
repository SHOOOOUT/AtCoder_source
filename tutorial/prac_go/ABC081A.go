package main

import (
	f "fmt"
	s "strings"
)

func main() {
	var a string

	f.Scanf("%s", &a)
	f.Printf("%d\n", s.Count(a, "1"))

}
