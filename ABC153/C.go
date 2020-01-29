package main

import (
	"bufio"
	f "fmt"
	"os"
	"sort"
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

	hp := scanSlice(n)
	sort.Ints(hp)

	flag := 0
	atack := 0

	if k >= n {
		f.Println(0)
	} else {
		for i := 0; i < n; i++ {
			flag++
			if flag > k {
				atack += hp[n-1-i]
			}
		}
		f.Println(atack)
	}

}
