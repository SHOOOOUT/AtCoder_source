package main

import(
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

func main(){
	sc.Split(bufio.ScanWords)
	num := scanSlice(5)

	for i, _ := range num {
		if num[i] == 0 {
			f.Println(i+1)
			break;
		}
	}
}