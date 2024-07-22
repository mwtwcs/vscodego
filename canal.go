package main

import (
	"fmt"
	"time"
)

func outd(start, end int) {
	for i := start; i <= end; i++ {
		fmt.Println(i)
	}
}

func canal() {
	// Запуск двух горутин
	go out(0, 5)
	go out(6, 10)

	// Ожидание завершения горутин
	time.Sleep(500 * time.Millisecond)
}