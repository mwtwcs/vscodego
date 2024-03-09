package main

import (
	"fmt"
)

func main() {
	var x1, x2 int
	fmt.Scan(&x1,&x2)

	if x1 % x2 == 0{
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}