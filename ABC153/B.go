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

	var hp, n int
	f.Scan(&hp, &n)

	skill_powers := scanSlice(n)
	sort.Ints(skill_powers)

	damage := 0

	for i := 0; i < n; i++ {
		damage += skill_powers[len(skill_powers)-1-i]
		num := i
		if damage >= hp {
			f.Println("Yes")
			break
		} else if num == n-1 && damage < hp {
			f.Println("No")
			break
		}
	}

}
