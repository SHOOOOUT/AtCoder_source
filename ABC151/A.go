package main

import (
	f "fmt"
)

func main() {
	a := [...]string{"a", "b", "c", "d", "e", "f", "g",
		"h", "i", "j", "k", "l", "n", "m", "o",
		"p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	var str string
	f.Scanf("%s", &str)

	for i := 0; i < 26; i++ {
		if a[i] == str {
			f.Printf("%s\n", a[i+1])
			break
		}
	}
}
