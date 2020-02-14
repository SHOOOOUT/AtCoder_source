package main

import (
	f "fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var s string
	f.Scan(&s)
	f.Println(strings.Repeat("x", utf8.RuneCountInString(s)))

}
