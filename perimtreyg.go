package main

import (
	"fmt"
	"math"
)
func perimetr() {
	var a, b float64
	fmt.Scan(&a,&b)
    summ := math.Sqrt(math.Pow(a,2) + math.Pow(b , 2))
	perimetr := a + b + summ
	fmt.Println(perimetr)
    

}