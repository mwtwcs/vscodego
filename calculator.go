package main

import "fmt"

func calculator() {
	var a, b float64
	var c string
	fmt.Scan(&a,&b,&c)
	if c == "+"{
		fmt.Print(a + b)
	} else if c == "-"{
		fmt.Print( a - b)
	} else if c == "*"{
		fmt.Print(a * b)
	} else if c == "/" {
		if b == 0{
			fmt.Print("На ноль делить нельзя!")
		} else {
			fmt.Print( a / b)
		}
    } else {
		fmt.Print("Неверная операция")
	}
}