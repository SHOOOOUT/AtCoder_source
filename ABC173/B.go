package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

/* 数字を1文字読み込む */
func scanStr() string {
	sc.Scan()
	a, err := sc.Text()
	if err != nil {
		fmt.Println(err)
	}
	return a
}

func scanSlice(n int) []string {
	a := make([]string, n)
	for i := 0; i < n; i++ {
		a[i] = scanStr()
	}
	return a
}

func main() {
	sc.Split(bufio.ScanWords)
	var n int
	fmt.Scan(&n)
	judge := scanSlice(n)
	fmt.Println(judge)
}
