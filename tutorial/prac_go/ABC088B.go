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

	num := scanInt()
	cards := scanSlice(num)

	sort.Ints(cards)

	var a, b int
	cnt := 0

	for i := num - 1; i >= 0; i-- {
		if cnt%2 == 0 {
			a += cards[i]
		} else if cnt%2 != 0 {
			b += cards[i]
		}
		cnt++
	}

	res := a - b

	f.Println(res)

}
