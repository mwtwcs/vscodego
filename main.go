package main

import (
	"fmt"
)

// Функция download() для имитации загрузки файла
func download(size int, ch chan int) {
	sum := 0
	for i := 0; i <= size; i++ {
		sum += i
	}
	ch <- sum // отправляем результат в канал
}

func main() {
	// Создаем каналы для передачи результатов
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	// Входные данные - размеры файлов
	s1 := 10
	s2 := 100
	s3 := 42

	// Запускаем горутины для каждого размера файла
	go download(s1, ch1)
	go download(s2, ch2)
	go download(s3, ch3)

	// Собираем результаты из каналов
	result1 := <-ch1
	result2 := <-ch2
	result3 := <-ch3

	// Вычисляем общую сумму результатов
	totalSum := result1 + result2 + result3

	// Выводим общую сумму на экран
	fmt.Println(totalSum)
}