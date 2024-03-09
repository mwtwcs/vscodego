package main

import (
	"fmt"
)

func number() {
	var str1, str2 int
	fmt.Scan(&str1)
	fmt.Println("Первое число:", str1)
	fmt.Scan(&str2)
	fmt.Println("Введите второе число:", str2)
	fmt.Println("Сумма:", str1+str2)
	fmt.Println("Разность:", str1-str2)
	fmt.Println("Произведение:", str1*str2)

	if str2 != 0 {
		fmt.Println("Частное:", float64(str1)/float64(str2))
	} else {
		fmt.Println("Деление на ноль невозможно.")
	}
}
