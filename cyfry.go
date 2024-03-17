package main

import "fmt"

func cyfry() {
	var a int
	fmt.Scan(&a)
	b := a / 100
	c := (a % 100) / 10
	d := a % 10
	if b != c && b != d && c != d{
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}