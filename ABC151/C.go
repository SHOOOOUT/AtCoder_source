/* 通せなかった */
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

/* 文字の読み込み */
func scanText() string {
	sc.Scan()
	return sc.Text()
}

func scanInt2() (int, int) {
	return scanInt(), scanInt()
}

func scan2() (int, string) {
	return scanInt(), scanText()
}

func main() {
	sc.Split(bufio.ScanWords)

	num, submission := scanInt2()
	f.Printf("%d %d\n", num, submission)

	array := make(map[int]string)

	for i := 0; i < submission; i++ {
		question, judge := scan2()
		array[i+question] = judge
	}

}
