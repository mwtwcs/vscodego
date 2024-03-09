package main

import "fmt"

func krug() {
	var r int
	var p float32
	fmt.Scan(&r)
	summ1 := r * r
	p = 3.14
	//s := summ1 * float64(p)
	fmt.Println(float32(p) * float32(summ1))
}
