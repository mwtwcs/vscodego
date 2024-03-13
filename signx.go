package main

import (
	"fmt"
)

func signx() {
	var x1 int
	fmt.Scan(&x1)
	
	if x1 < 0 {
		fmt.Print("-1")
	}
	if x1 == 0 {
        fmt.Print("0")
	}
	if x1 > 1 {
		fmt.Print("1")
	}
}
