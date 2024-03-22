package main

import "fmt"

func simmetrica() {
	var a string
	fmt.Scan(&a)
	if a[0] == a[3] && a[1] == a[2] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}