package main

import "fmt"

func visokos() {
	var year int
	fmt.Scan(&year)
	if (year%4 == 0 && year%100 != 0) || year%400 == 0{
		fmt.Print("YES")
    } else {
		fmt.Print("NO")
	} 
}
