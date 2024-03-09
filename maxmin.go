package main

import (
	"fmt"
	
)

func maxmin(){
    var x1, x2 float64
    fmt.Scan(&x1, &x2)

	max := x1
	if x2 > x1 {
		max = x2
	}
	fmt.Println(max)
}