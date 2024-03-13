package main

import "fmt"

func defera() {
	defer fmt.Println("1") //отложенное действие
	defer fmt.Println("2") //отложенное действие
	defer fmt.Println("3") //отложенное действие
	defer fmt.Println("4") //отложенное действие
	fmt.Println("hello")
}

// stack, Last in First Out выполняется последниее из первого