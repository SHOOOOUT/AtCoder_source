package main

import (
	f "fmt"
	"strconv"
	"strings"
)

func main() {

	var a, b int
	f.Scan(&a, &b)

	i := strings.Repeat(strconv.Itoa(a), b)
	j := strings.Repeat(strconv.Itoa(b), a)

	if i <= j {
		f.Println(i)
	} else if j < i {
		f.Println(j)
	}

}
