package main

import (
	f "fmt"
)

func main() {
	var x int
	f.Scan(&x)

	if 400<=x && x<=599{
		f.Println(8)
	}else if 600<=x && x<=799{
		f.Println(7)
	}else if 800<=x && x<=999{
		f.Println(6)
	}else if 1000<=x && x<=1199{
		f.Println(5)
	}else if 1200<=x && x<=1399{
		f.Println(4)
	}else if 1400<=x && x<=1599{
		f.Println(3)
	}else if 1600<=x && x<=1799{
		f.Println(2)
	}else if 1800<=x && x<=1999{
		f.Println(1)
	}
}