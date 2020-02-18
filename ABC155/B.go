package main

import (
	"bufio"
	f "fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	a, err := strconv.Atoi(sc.Text())
	if err != nil {
		f.Println(err)
	}
	return a
}

func scanSlice(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}
	return a
}

func main() {
	sc.Split(bufio.ScanWords)

	n := scanInt()
	num := scanSlice(n)

	flag := 0

	for i := 0; i < n; i++ {
		if num[i]%2 == 0 && num[i]%3 != 0 && num[i]%5 != 0 {
			flag++
			f.Println("DENIED")
			break
		}
	}

	if flag == 0 {
		f.Println("APPROVED")
	}
}
