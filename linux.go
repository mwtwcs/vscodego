package main

import "fmt"

func linux() {
	var a int
	fmt.Scan(&a)
	b := a / 1000
	c := a % 1000
	one := (b / 100) + ((b % 100) / 10) + (b % 10)
	two := (c / 100) + ((c % 100) / 10) + (c % 10)
	if one == two {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
