package main

import (
	"fmt"
)

func chet() {
	var x1 int
	fmt.Scan(&x1)
	if x1 % 2 == 0 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
