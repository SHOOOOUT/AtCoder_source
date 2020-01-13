package main

import (
	f "fmt"
)

func main() {

	var a, b, c int
	var sum int
	f.Scan(&a, &b, &c, &sum)

	cnt := 0 //何通りのパターンがあったか

	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			for k := 0; k <= c; k++ {
				if i*500+j*100+k*50 == sum {
					cnt++
				}
			}
		}
	}

	f.Println(cnt)

}
