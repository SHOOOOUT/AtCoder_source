package main

import (
	"bufio"
	f "fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

/* 数字を1文字読み込む */
func scanInt() int {
	sc.Scan()
	a, err := strconv.Atoi(sc.Text())
	if err != nil {
		f.Println(err)
	}
	return a
}

/* 動的にスライスに格納 */
func scanSlice(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}
	return a
}

func main() {
	sc.Split(bufio.ScanWords)
	var n int
	f.Scan(&n)
	arr := scanSlice(n)

	count := 0

	//ここのforの回し方がわからない
	for i := 0; i < n; {
		i++
		for j := 0; j < i; {
			j++
			for k := 0; k < j; {
				k++
				if arr[i] != arr[j] && arr[i] != arr[j] && arr[k]+arr[j] > arr[i] {
					count++
				}
			}
		}
	}
	f.Println(count)
}
