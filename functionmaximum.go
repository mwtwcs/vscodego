package main

import (
	"fmt"
)
func maximum(a int, b int) int{
	if a > b{
		return a
	} else{
		return b
	}
}

func funcionmaximum(){
	fmt.Println(maximum(10,11))
}