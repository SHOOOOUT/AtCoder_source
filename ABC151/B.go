package main

import (
	"bufio"
	"fmt"
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
		fmt.Println(err)
	}
	return a
}

/* 複数文字連続で取得するとき */
func scanInt3() (int, int, int) {
	return scanInt(), scanInt(), scanInt()
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

	num, max, ave := scanInt3()

	score := scanSlice(num - 1)
	score_sum := 0

	for i := 0; i < num-1; i++ {
		score_sum += score[i]
	}

	mark := (ave * num) - score_sum
	failure := (score_sum + max) / num

	if mark <= max && score_sum/num < ave {
		f.Println(mark)
	} else if score_sum/num >= ave {
		f.Println("0")
	} else if failure < ave {
		f.Println(-1)
	}

}
