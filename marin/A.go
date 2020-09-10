package main

import (
	f "fmt"
)

func main() {
	var a string

	f.Scan(&a)
	f.Println(a[0:3])
}
