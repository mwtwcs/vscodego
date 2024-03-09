package main

import (
	"bufio"
	"fmt"
	"os"
)

func output() {
	scanner := bufio.NewScanner(os.Stdin) // создаем копию структуры bufio.Scanner
	_ = scanner.Scan()                    // на этом месте приложение останавливается и ожидает ввода. Завершением ввода, будет нажатие Enter
	name := scanner.Text()                // cохраняем все, что ввели в переменную "name"
	// Формируем выходную строку
	output := name + " - лучшая книга!"
	// Выводим результат
	fmt.Println(output)
}
