package main

import "fmt"

func stepen() {
	var stepen int
	fmt.Scan(&stepen)
	stepen1 := stepen * stepen   // x^2
	stepen3 := stepen1 * stepen1 // x^4
	stepen5 := stepen3 * stepen1 // x^6
	fmt.Println(stepen5)
}
