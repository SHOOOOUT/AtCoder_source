package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	arr := strings.Split(s, "")

	if arr[0] == "R" && arr[1] == "R" && arr[2] == "R" {
		fmt.Println(3)
	}
	if (arr[0] == "R" && arr[1] == "R" && arr[2] == "S") || arr[0] == "S" && arr[1] == "R" && arr[2] == "R" {
		fmt.Println(2)
	}
	if (arr[0] == "R" && arr[1] == "S" && arr[2] == "S") || arr[0] == "R" && arr[1] == "S" && arr[2] == "R" || arr[0] == "S" && arr[1] == "S" && arr[2] == "R" || arr[0] == "S" && arr[1] == "R" && arr[2] == "S" {
		fmt.Println(1)
	}
	if arr[0] == "S" && arr[1] == "S" && arr[2] == "S" {
		fmt.Println(0)
	}
}
