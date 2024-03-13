package main

import "fmt"

func lessons1() {
	a := 0
	for a < 1000{
		if a == 100{
			fmt.Println(a ,"is full ")
		} else {
			fmt.Println(fmt.Sprintf(" a is not 100 a=%d", a)) // вывод полной консоли воспроизведения кода
		}
		a ++
	}
}