package main

import (
	"bufio"
	f "fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanSlice(n int) ([]int, []int) {
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i] = scanInt2()
	}
	return a, b
}

/* 数字を1文字読み込む */
func scanInt() int {
	sc.Scan()
	a, err := strconv.Atoi(sc.Text())
	if err != nil {
		f.Println(err)
	}
	return a
}

/* 複数文字連続で取得するとき */
func scanInt2() (int, int) {
	return scanInt(), scanInt()
}

func main() {
	sc.Split(bufio.ScanWords)

	n, d := scanInt2()
	dFloat := float64(d)

	x, y := scanSlice(n)

	cnt := 0
	for i := 0; i < n; i++ {
		xF := float64(x[i])
		yF := float64(y[i])
		sum := math.Sqrt(xF*xF + yF*yF)
		if sum <= dFloat {
			cnt += 1
		}
	}

	f.Println(cnt)
}
