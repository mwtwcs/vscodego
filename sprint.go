package main

import "fmt"

func sprint() {
	var name, age = "Arthur", 32
	var c = fmt.Sprintf("My name is %s. And im %d years old.", name, age)
	fmt.Println(c)
}
