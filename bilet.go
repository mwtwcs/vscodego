package main

import "fmt"

func bilet() {
	var a int
	fmt.Scan(&a)
	b := a / 1000
	c := a % 1000
	one := (b / 100) + ((b % 100) / 10) + (b % 10)
	two := (c / 100) + ((c % 100) / 10) + (c % 10)
	if one == two {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}

//package main

//import "fmt"

//func main() {
//	var num string
//fmt.Scan(&num)
//	if num[0]+num[1]+num[2] == num[3]+num[4]+num[5] {
//		fmt.Println("YES")
//		return
//	}
//	fmt.Println("NO")
//}
