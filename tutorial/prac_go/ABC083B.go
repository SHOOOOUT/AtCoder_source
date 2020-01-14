package main

import (
	f "fmt"
)

func sum(x int) int {
	if x == 0 {
		return 0
	}
	return sum(x/10) + x%10
}

func main() {

	var x, a, b int
	f.Scan(&x, &a, &b)

	res := 0

	for i := 1; i <= x; i++ {
		ans := sum(i)
		if a <= ans && ans <= b {
			res += i
		}
	}

	f.Println(res)

}
