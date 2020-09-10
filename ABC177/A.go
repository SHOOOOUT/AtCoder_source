package main

import (
	f "fmt"
)

func main() {
	var d, t, s int
	f.Scan(&d, &t, &s)

	if s*t >= d {
		f.Println("Yes")
	} else {
		f.Println("No")
	}
}
