package main

import (
	"fmt"
	"math"
)

func manhetten() {
	var x1, x2, y1, y2 float64
	fmt.Scan(&x1,&y1,&x2,&y2)
	summ := math.Abs(x1-x2) + math.Abs(y1-y2)
	fmt.Println(summ)
}