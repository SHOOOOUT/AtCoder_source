package main

import (
	f "fmt"
)

func main() {
	var n int
	f.Scan(&n)

	var diff int
	for i := 1; ; i++ {
		diff = 1000*i - n
		if diff > 0 && diff < 1000 {
			f.Println(diff)
			break
		} else if diff == 0 {
			f.Println(0)
			break
		}
	}
}
