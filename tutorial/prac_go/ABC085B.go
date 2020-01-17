package main

import (
	f "fmt"
)

func main() {

	var num int
	f.Scan(&num)
	array := map[int]int{}

	//arrayのkeyにtmpを指定する
	//その長さを求めれば良い
	for i := 0; i < num; i++ {
		var tmp int
		f.Scan(&tmp)
		array[tmp] = 1
	}

	f.Println(len(array))

}
