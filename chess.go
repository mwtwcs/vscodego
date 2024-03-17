package main

import "fmt"

func chess() {
	var a1 , a2 , b1, b2 int
	fmt.Scan(&a1,&a2,&b1,&b2)
	if a1 == b1 || a2 == b2{
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	} 
}