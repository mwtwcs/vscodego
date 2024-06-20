package main

import "fmt"


func squareSumm(a int, b int)int{
	sum := 0
	for i := a; i <= b; i++{
		sum+= i*i
	}
	return sum
}



func cyclesquase(){
	fmt.Println(squareSumm(2,6))
}