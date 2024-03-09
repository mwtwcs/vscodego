package main

import "fmt"

func code() {
	var name string
	age := 5
	fmt.Println("What is your name")
	fmt.Scan(&name)
	fmt.Println("Hello " + name + "!")
	fmt.Println("How old are you")
	fmt.Scan(&age)
	fmt.Println("You are " + fmt.Sprint(age) + " years!")
}
