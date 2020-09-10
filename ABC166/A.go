package main

import (
	f "fmt"
)

func main(){
	var a string
	f.Scan(&a)

	if a == "ABC"{
		f.Println("ARC")
	}else if a == "ARC"{
		f.Println("ABC")
	}
}