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

	var n, k int
	f.Scan(&n, &k)
	a := scanSlice(n)

	num := make([]int, n)
	for i:=k; i<=n; i++{
		s := a[i] * a[i-1] * a[i-2]
		num := append(num, s)
		f.Println(num)
	}


	f.Println(a)

}
