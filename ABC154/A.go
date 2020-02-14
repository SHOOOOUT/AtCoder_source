package main

import (
	f "fmt"
)

func main() {

	var s, t string
	f.Scan(&s, &t)

	var s_num, t_num int
	f.Scan(&s_num, &t_num)

	var line string
	f.Scan(&line)

	if line == s {
		f.Println(s_num-1, t_num)
	} else if line == t {
		f.Println(s_num, t_num-1)
	}

}
