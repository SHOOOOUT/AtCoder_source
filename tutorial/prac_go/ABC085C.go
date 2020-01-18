package main

import (
	f "fmt"
)

func main() {

	var num int
	f.Scan(&num)
	var sum int
	f.Scan(&sum)

	var x, y int
	for x = 0; x <= num; x++ {
		for y = 0; y <= num-x; y++ {
			if 10000*x+5000*y+1000*(num-x-y) == sum {
				f.Printf("%d %d %d", x, y, num-x-y)
				return
			}
		}
	}

	f.Println("-1 -1 -1")

}
