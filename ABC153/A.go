package main

import (
	f "fmt"
)

func main() {
	var hp, a int
	f.Scan(&hp, &a)

	sum := 0
	num := 0

	for i := 1; ; i++ {
		sum += a
		num = i
		if sum >= hp {
			break
		}
	}

	f.Println(num)
}
