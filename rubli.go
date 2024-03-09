package main

import (
	"fmt"
)

func rubli() {
	// Ввод данных
	var a, b, n int
	fmt.Scan(&a, &b, &n)

	// Вычисление общей стоимости в копейках
	totalInKopecks := (a*100 + b) * n

	// Пересчет в рубли и копейки
	rub := totalInKopecks / 100
	kop := totalInKopecks % 100

	// Вывод результатов
	fmt.Printf("%d %d\n", rub, kop)
}
